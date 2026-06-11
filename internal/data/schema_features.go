package data

import (
	"fmt"
	"strings"

	"github.com/project-kessel/inventory-api/internal/biz/model"
	"github.com/xeipuuv/gojsonschema"
)

// NewSchemaFactoryWithFeatures returns a SchemaFromString factory that dispatches to
// FeaturesServiceSchema when the JSON schema contains "allowed_workspace_ids",
// otherwise falls back to the standard JsonSchemaWithWorkspaces.
func NewSchemaFactoryWithFeatures() model.SchemaFromString {
	return func(jsonSchema string) model.Schema {
		if strings.Contains(jsonSchema, "allowed_workspace_ids") {
			return NewFeaturesServiceSchemaFromString(jsonSchema)
		}
		return NewJsonSchemaWithWorkspacesFromString(jsonSchema)
	}
}

// FeaturesServiceSchema handles tuple calculation for "service" resources reported by the
// "features" reporter. It creates allowed_workspaces, billing_account, and parent tuples
// in addition to the standard workspace tuple.
type FeaturesServiceSchema struct {
	jsonSchema string
}

func NewFeaturesServiceSchemaFromString(jsonSchema string) model.Schema {
	return FeaturesServiceSchema{jsonSchema: jsonSchema}
}

func (s FeaturesServiceSchema) Validate(data interface{}) (bool, error) {
	schemaLoader := gojsonschema.NewStringLoader(s.jsonSchema)
	dataLoader := gojsonschema.NewGoLoader(data)
	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		return false, fmt.Errorf("validation error: %w", err)
	}
	if !result.Valid() {
		var errMsgs []string
		for _, desc := range result.Errors() {
			errMsgs = append(errMsgs, desc.String())
		}
		return false, fmt.Errorf("validation failed: %s", strings.Join(errMsgs, "; "))
	}
	return true, nil
}

func (s FeaturesServiceSchema) CalculateTuples(current, previous *model.Representations, key model.ReporterResourceKey) (model.TuplesToReplicate, error) {
	var tuplesToCreate, tuplesToDelete []model.RelationsTuple

	currentAllowedWS := getStringSlice(current)
	previousAllowedWS := getStringSlice(previous)
	currentBA := getBillingAccount(current)
	previousBA := getBillingAccount(previous)
	currentParent := getParentService(current)
	previousParent := getParentService(previous)

	// allowed_workspaces: create tuples for each workspace in the list
	added, removed := diffStringSlices(previousAllowedWS, currentAllowedWS)
	for _, ws := range added {
		tuplesToCreate = append(tuplesToCreate, model.NewAllowedWorkspacesTuple(ws, key))
	}
	for _, ws := range removed {
		tuplesToDelete = append(tuplesToDelete, model.NewAllowedWorkspacesTuple(ws, key))
	}

	// billing_account relation
	if currentBA != previousBA {
		if currentBA != "" {
			tuplesToCreate = append(tuplesToCreate, model.NewBillingAccountTuple(currentBA, key))
		}
		if previousBA != "" {
			tuplesToDelete = append(tuplesToDelete, model.NewBillingAccountTuple(previousBA, key))
		}
	}

	// parent service relation
	if currentParent != previousParent {
		if currentParent != "" {
			tuplesToCreate = append(tuplesToCreate, model.NewParentServiceTuple(currentParent, key))
		}
		if previousParent != "" {
			tuplesToDelete = append(tuplesToDelete, model.NewParentServiceTuple(previousParent, key))
		}
	}

	return model.NewTuplesToReplicate(tuplesToCreate, tuplesToDelete)
}

func getStringSlice(r *model.Representations) []string {
	if r == nil {
		return nil
	}
	return r.AllowedWorkspaceIDs()
}

func getBillingAccount(r *model.Representations) string {
	if r == nil {
		return ""
	}
	return r.BillingAccountID()
}

func getParentService(r *model.Representations) string {
	if r == nil {
		return ""
	}
	return r.ParentServiceID()
}

// diffStringSlices returns elements added to current and removed from previous.
func diffStringSlices(previous, current []string) (added, removed []string) {
	prevSet := make(map[string]struct{}, len(previous))
	for _, s := range previous {
		prevSet[s] = struct{}{}
	}
	currSet := make(map[string]struct{}, len(current))
	for _, s := range current {
		currSet[s] = struct{}{}
	}

	for _, s := range current {
		if _, ok := prevSet[s]; !ok {
			added = append(added, s)
		}
	}
	for _, s := range previous {
		if _, ok := currSet[s]; !ok {
			removed = append(removed, s)
		}
	}
	return
}
