package model_legacy

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/project-kessel/inventory-api/internal"
	datamodel "github.com/project-kessel/inventory-api/internal/data/model"
	kessel "github.com/project-kessel/relations-api/api/kessel/relations/v1beta1"
	"google.golang.org/protobuf/proto"
)

func newResourceEventLegacy(operationType internal.EventOperationType, resource *Resource) (*datamodel.ResourceEvent, error) {
	const eventType = "resources"
	now := time.Now()

	eventId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	var labels []datamodel.EventResourceLabel
	for _, val := range resource.Labels {
		labels = append(labels, datamodel.EventResourceLabel(val))
	}

	var reportedTime time.Time
	var createdAt *time.Time
	var updatedAt *time.Time
	var deletedAt *time.Time

	switch operationType {
	case internal.OperationTypeCreated:
		createdAt = resource.CreatedAt
		reportedTime = *createdAt
	case internal.OperationTypeUpdated:
		updatedAt = resource.UpdatedAt
		reportedTime = *updatedAt
	case internal.OperationTypeDeleted:
		deletedAt = &now
		reportedTime = *deletedAt
	}

	return &datamodel.ResourceEvent{
		Specversion:     "1.0",
		Type:            makeEventType(eventType, resource.ResourceType, string(operationType.OperationType())),
		Source:          "", // TODO: inventory uri
		Id:              eventId.String(),
		Subject:         makeEventSubject(eventType, resource.ResourceType, resource.ID.String()),
		Time:            reportedTime,
		DataContentType: "application/json",
		Data: datamodel.EventResourceData{
			Metadata: datamodel.EventResourceMetadata{
				Id:           resource.ID.String(),
				OrgId:        resource.OrgId,
				ResourceType: resource.ResourceType,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
				DeletedAt:    deletedAt,
				WorkspaceId:  resource.WorkspaceId,
				Labels:       labels,
			},
			ReporterData: datamodel.EventResourceReporter{
				ReporterInstanceId: resource.ReporterId,
				ReporterType:       resource.Reporter.ReporterType, //nolint:staticcheck
				ConsoleHref:        resource.ConsoleHref,
				ApiHref:            resource.ApiHref,
				LocalResourceId:    resource.ReporterResourceId,
				ReporterVersion:    &resource.Reporter.ReporterVersion, //nolint:staticcheck
			},
			ResourceData: resource.ResourceData,
		},
	}, nil
}

func convertResourceToResourceEventLegacy(resource Resource, operationType internal.EventOperationType) (internal.JsonObject, error) {
	payload := internal.JsonObject{}

	resourceEvent, err := newResourceEventLegacy(operationType, &resource)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource event: %w", err)
	}

	marshalledJson, err := json.Marshal(resourceEvent)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource to json: %w", err)
	}

	err = json.Unmarshal(marshalledJson, &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to payload: %w", err)
	}

	return payload, nil
}

func convertResourceToSetTupleEventLegacy(resource Resource, namespace string) (internal.JsonObject, error) {
	payload := internal.JsonObject{}

	// Derive namespace from the resource when possible
	if resource.ReporterType != "" {
		namespace = strings.ToLower(resource.ReporterType)
	} else if resource.Reporter.ReporterType != "" { //nolint:staticcheck
		namespace = strings.ToLower(resource.Reporter.ReporterType)
	}

	relationship := &kessel.Relationship{
		Resource: &kessel.ObjectReference{
			Type: &kessel.ObjectType{
				Name:      resource.ResourceType,
				Namespace: namespace,
			},
			Id: resource.ReporterResourceId,
		},
		Relation: "workspace",
		Subject: &kessel.SubjectReference{
			Subject: &kessel.ObjectReference{
				Type: &kessel.ObjectType{
					Name:      "workspace",
					Namespace: "rbac",
				},
				Id: resource.WorkspaceId,
			},
		},
	}

	marshalledJson, err := json.Marshal(relationship)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource to json: %w", err)
	}
	err = json.Unmarshal(marshalledJson, &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to payload: %w", err)
	}

	return payload, nil
}

func convertResourceToUnsetTupleEventLegacy(resource Resource, namespace string) (internal.JsonObject, error) {
	payload := internal.JsonObject{}

	// Derive namespace from the resource when possible
	if resource.ReporterType != "" {
		namespace = strings.ToLower(resource.ReporterType)
	} else if resource.Reporter.ReporterType != "" { //nolint:staticcheck
		namespace = strings.ToLower(resource.Reporter.ReporterType)
	}

	tuple := &kessel.RelationTupleFilter{
		ResourceNamespace: proto.String(namespace),
		ResourceType:      proto.String(resource.ResourceType),
		ResourceId:        proto.String(resource.ReporterResourceId),
		Relation:          proto.String("workspace"),
	}

	marshalledJson, err := json.Marshal(tuple)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource to json: %w", err)
	}
	err = json.Unmarshal(marshalledJson, &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to payload: %w", err)
	}

	return payload, nil
}

func NewOutboxEventsFromResource(resource Resource, namespace string, operationType internal.EventOperationType, txid string) (*datamodel.OutboxEvent, *datamodel.OutboxEvent, error) {
	var tuplePayload internal.JsonObject
	var tupleEvent *datamodel.OutboxEvent

	// Build resource event
	payload, err := convertResourceToResourceEventLegacy(resource, operationType)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert resource to payload: %w", err)
	}

	resourceEvent := &datamodel.OutboxEvent{
		Operation:     operationType,
		AggregateType: datamodel.ResourceAggregateType,
		AggregateID:   resource.InventoryId.String(),
		TxId:          "",
		Payload:       payload,
	}

	// Build tuple event
	switch operationType.OperationType() {
	case internal.OperationTypeDeleted:
		tuplePayload, err = convertResourceToUnsetTupleEventLegacy(resource, namespace)
	default:
		tuplePayload, err = convertResourceToSetTupleEventLegacy(resource, namespace)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert resource to payload: %w", err)
	}

	tupleEvent = &datamodel.OutboxEvent{
		Operation:     operationType,
		AggregateType: datamodel.TupleAggregateType,
		AggregateID:   resource.InventoryId.String(),
		TxId:          txid,
		Payload:       tuplePayload,
	}

	return resourceEvent, tupleEvent, nil
}

func makeEventType(eventType, resourceType, operation string) string {
	return fmt.Sprintf("redhat.inventory.%s.%s.%s", eventType, resourceType, operation)
}

func makeEventSubject(eventType, resourceType, resourceId string) string {
	return "/" + strings.Join([]string{eventType, resourceType, resourceId}, "/")
}
