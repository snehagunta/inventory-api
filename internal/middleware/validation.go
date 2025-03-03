package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	rel "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/relationships"
	pb "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	pb2 "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta2/resources"
	"github.com/xeipuuv/gojsonschema"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
	"path/filepath"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/protobuf/proto"
)

var (
	resourceDir       = os.Getenv("RESOURCE_DIR")
	AbstractResources = map[string]struct{}{} // Tracks resource types marked as abstract (no resource_data)
)

func isDeleteRequest(v interface{}) bool {
	switch v.(type) {
	case *pb.DeleteK8SClusterRequest,
		*pb.DeleteRhelHostRequest,
		*pb.DeleteK8SPolicyRequest,
		*pb.DeleteNotificationsIntegrationRequest,
		*pb2.DeleteResourceRequest:
		return true
	default:
		return false
	}
}

func isRelationshipRequest(v interface{}) bool {
	switch v.(type) {
	case *rel.CreateK8SPolicyIsPropagatedToK8SClusterRequest,
		*rel.UpdateK8SPolicyIsPropagatedToK8SClusterRequest,
		*rel.DeleteK8SPolicyIsPropagatedToK8SClusterRequest:
		return true
	default:
		return false
	}
}

func Validation(validator protovalidate.Validator) middleware.Middleware {
	if resourceDirFilePath, exists := os.LookupEnv("RESOURCE_DIR"); exists {
		fmt.Println(resourceDirFilePath)
		absPath, err := filepath.Abs(resourceDirFilePath)
		if err != nil {
			log.Errorf("failed to resolve absolute path for RESOURCE_DIR file: %v", err)
		}
		resourceDir = absPath
	}

	if err := PreloadAllSchemas(resourceDir); err != nil {
		log.Fatalf("Failed to preload schemas: %v", err)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if v, ok := req.(proto.Message); ok {
				if isDeleteRequest(v) || isRelationshipRequest(v) {
					// run the protovalidate validation if it is a delete or relationship request
					if err := validator.Validate(v); err != nil {
						return nil, errors.BadRequest("VALIDATOR", err.Error()).WithCause(err)
					}
				} else {
					// Otherwise, run both protovalidate and JSON validation.
					if err := validator.Validate(v); err != nil {
						return nil, errors.BadRequest("VALIDATOR", err.Error()).WithCause(err)
					}
					if err := validateResourceJSON(v); err != nil {
						return nil, errors.BadRequest("JSON_VALIDATOR", err.Error()).WithCause(err)
					}
				}
			}
			return handler(ctx, req)
		}
	}
}

func validateResourceJSON(msg proto.Message) error {
	fmt.Println("DEBUG: Starting JSON validation...")

	// Step 1: Marshal proto message to JSON
	data, err := protojson.Marshal(msg)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to marshal message: %w", err)
	}
	fmt.Printf("DEBUG: Raw JSON: %s\n", string(data))

	// Step 2: Unmarshal JSON into a map
	var resourceMap map[string]interface{}
	if err := json.Unmarshal(data, &resourceMap); err != nil {
		return fmt.Errorf("ERROR: Failed to unmarshal JSON: %w", err)
	}
	fmt.Printf("DEBUG: Parsed resourceMap: %+v\n", resourceMap)

	// Step 3: Extract `resource` field
	resource, ok := resourceMap["resource"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("ERROR: Missing or invalid 'resource' field in payload")
	}
	fmt.Printf("DEBUG: Extracted resource: %+v\n", resource)

	// Step 4: Extract `resourceType`
	resourceTypeRaw, exists := resource["resourceType"]
	if !exists {
		return fmt.Errorf("ERROR: Missing 'resourceType' field in resource payload")
	}
	fmt.Printf("DEBUG: Raw resourceType value: %#v\n", resourceTypeRaw)

	resourceType, ok := resourceTypeRaw.(string)
	if !ok || resourceType == "" {
		return fmt.Errorf("ERROR: 'resourceType' is not a valid string (got %T instead)", resourceTypeRaw)
	}
	fmt.Printf("DEBUG: Extracted resourceType: %s\n", resourceType)

	// Step 5: Extract and validate `reporterData`
	reporterData, ok := resource["reporterData"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("ERROR: Missing or invalid 'reporterData' field for resource '%s'", resourceType)
	}
	fmt.Printf("DEBUG: Extracted reporterData: %+v\n", reporterData)

	reporterType, ok := reporterData["reporterType"].(string)
	if !ok {
		return fmt.Errorf("ERROR: Missing or invalid 'reporterType' field for resource '%s'", resourceType)
	}
	fmt.Printf("DEBUG: Extracted reporterType: %s\n", reporterType)

	// Step 6: Ensure `resourceData` is properly parsed
	resourceData, exists := reporterData["resourceData"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("missing or invalid reporter_data field for resource '%s'", resourceData)
	}

	// Step 7: Ensure `commonResourceData` is properly parsed
	commonResourceData, exists := resource["commonResourceData"]
	if exists {
		switch v := commonResourceData.(type) {
		case string:
			var parsedData map[string]interface{}
			if err := json.Unmarshal([]byte(v), &parsedData); err != nil {
				return fmt.Errorf("ERROR: Failed to parse 'commonResourceData' JSON string: %w", err)
			}
			resource["commonResourceData"] = parsedData
			fmt.Printf("DEBUG: Parsed commonResourceData from string: %+v\n", parsedData)
		case map[string]interface{}:
			fmt.Printf("DEBUG: commonResourceData is already a structured object: %+v\n", v)
		default:
			return fmt.Errorf("ERROR: Unexpected 'commonResourceData' type: %T", v)
		}
	}

	// Step 8: Validate JSON schema fields
	if resourceData, resourceDataExists := reporterData["resourceData"].(map[string]interface{}); resourceDataExists {
		fmt.Println("DEBUG: Validating 'resourceData' against schema...")
		resourceSchema, err := getSchemaFromCache(fmt.Sprintf("%s:%s", "resource", strings.ToLower(resourceType)))
		if err != nil {
			return fmt.Errorf("ERROR: Failed to load resource schema for '%s': %w", resourceType, err)
		}
		if err := validateJSONSchema(resourceSchema, resourceData); err != nil {
			return fmt.Errorf("ERROR: 'resourceData' validation failed for '%s': %w", resourceType, err)
		}
	}

	if commonResourceData, commonResourceExists := resource["commonResourceData"].(map[string]interface{}); commonResourceExists {
		fmt.Println("DEBUG: Validating 'commonResourceData' against schema...")
		commonSchema, err := getSchemaFromCache(fmt.Sprintf("%s:%s", "common", strings.ToLower(resourceType)))
		if err != nil {
			return fmt.Errorf("ERROR: Failed to load common resource schema for '%s': %w", resourceType, err)
		}
		if err := validateJSONSchema(commonSchema, commonResourceData); err != nil {
			return fmt.Errorf("ERROR: 'commonResourceData' validation failed for '%s': %w", resourceType, err)
		}
	}

	//// Step 9: Validate resource-reporter compatibility
	//fmt.Println("DEBUG: Checking resource-reporter compatibility...")
	//if err := ValidateCombination(resourceType, reporterType); err != nil {
	//	return fmt.Errorf("ERROR: Resource-reporter compatibility validation failed for '%s': %w", resourceType, err)
	//}
	//
	//// Step 10: Validate `reporterData` using schema
	//fmt.Println("DEBUG: Validating 'reporterData' against schema...")
	//reporterSchema, err := getSchemaFromCache(fmt.Sprintf("%s:%s", resourceType, strings.ToLower(reporterType)))
	//if err != nil {
	//	return fmt.Errorf("ERROR: Failed to load reporter schema for '%s': %w", reporterType, err)
	//}
	//
	//if err := validateJSONSchema(reporterSchema, reporterData); err != nil {
	//	return fmt.Errorf("ERROR: 'reporterData' validation failed for resource '%s': %w", resourceType, err)
	//}

	fmt.Println("DEBUG: Validation successfully passed!")
	return nil
}

func validateJSONSchema(schemaStr string, jsonData interface{}) error {
	schemaLoader := gojsonschema.NewStringLoader(schemaStr)
	dataLoader := gojsonschema.NewGoLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	if !result.Valid() {
		var errMsgs []string
		for _, desc := range result.Errors() {
			errMsgs = append(errMsgs, desc.String())
		}
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, "; "))
	}
	return nil
}
