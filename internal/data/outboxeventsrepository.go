package data

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/project-kessel/inventory-api/internal"
	bizmodel "github.com/project-kessel/inventory-api/internal/biz/model"
	datamodel "github.com/project-kessel/inventory-api/internal/data/model"
	"gorm.io/gorm"
)

// publishes an event to the outbox table and then deletes it
// keeping the event in the outbox table is unnecessary
// as CDC will read from the write-ahead log
func PublishOutboxEvent(tx *gorm.DB, event *datamodel.OutboxEvent) error {
	if err := tx.Create(event).Error; err != nil {
		return err
	}
	if err := tx.Delete(event).Error; err != nil {
		return err
	}
	return nil
}

func NewOutboxEventsFromResourceEvent(domainResourceEvent bizmodel.ResourceReportEvent, operationType internal.EventOperationType, txid string) (*datamodel.OutboxEvent, *datamodel.OutboxEvent, error) {
	var tuplePayload internal.JsonObject
	var tupleEvent *datamodel.OutboxEvent

	// Build resource event
	payload, err := convertResourceToResourceEvent(domainResourceEvent, operationType)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert resource to payload: %w", err)
	}

	resourceEvent := &datamodel.OutboxEvent{
		Operation:     operationType,
		AggregateType: datamodel.ResourceAggregateType,
		AggregateID:   domainResourceEvent.Id().String(),
		TxId:          "",
		Payload:       payload,
	}

	// Build tuple event
	/*switch operationType.OperationType() {
	case OperationTypeDeleted:
		tuplePayload, err = convertResourceToUnsetTupleEvent(domainResourceEvent)
	default:
		tuplePayload, err = convertResourceToSetTupleEvent(domainResourceEvent)
	}*/

	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert resource to payload: %w", err)
	}

	tupleEvent = &datamodel.OutboxEvent{
		Operation:     operationType,
		AggregateType: datamodel.TupleAggregateType,
		AggregateID:   domainResourceEvent.Id().String(),
		TxId:          txid,
		Payload:       tuplePayload,
	}

	return resourceEvent, tupleEvent, nil
}

func convertResourceToResourceEvent(resourceReportEvent bizmodel.ResourceReportEvent, operationType internal.EventOperationType) (internal.JsonObject, error) {
	payload := internal.JsonObject{}

	resourceEvent, err := newResourceEvent(operationType, &resourceReportEvent)
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

func newResourceEvent(operationType internal.EventOperationType, resourceEvent *bizmodel.ResourceReportEvent) (*datamodel.ResourceEvent, error) {
	const eventType = "resources"
	now := time.Now()

	eventId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	var reportedTime time.Time
	var createdAt *time.Time
	var updatedAt *time.Time
	var deletedAt *time.Time

	switch operationType {
	case internal.OperationTypeCreated:
		createdAt = resourceEvent.CreatedAt()
		reportedTime = *createdAt
	case internal.OperationTypeUpdated:
		updatedAt = resourceEvent.UpdatedAt()
		reportedTime = *updatedAt
	case internal.OperationTypeDeleted:
		deletedAt = &now
		reportedTime = *deletedAt
	}

	return &datamodel.ResourceEvent{
		Specversion:     "1.0",
		Type:            makeEventType(eventType, resourceEvent.ResourceType(), string(operationType.OperationType())),
		Source:          "", // TODO: inventory uri
		Id:              eventId.String(),
		Subject:         makeEventSubject(eventType, resourceEvent.ResourceType(), resourceEvent.Id().String()),
		Time:            reportedTime,
		DataContentType: "application/json",
		Data: datamodel.EventResourceData{
			Metadata: datamodel.EventResourceMetadata{
				Id:           resourceEvent.Id().String(),
				ResourceType: resourceEvent.ResourceType(),
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
				DeletedAt:    deletedAt,
				WorkspaceId:  resourceEvent.WorkspaceId(),
			},
			ReporterData: datamodel.EventResourceReporter{
				ReporterInstanceId: resourceEvent.ReporterInstanceId(),
				ReporterType:       resourceEvent.ReporterType(),
				ConsoleHref:        resourceEvent.ConsoleHref(),
				ApiHref:            resourceEvent.ApiHref(),
				LocalResourceId:    resourceEvent.LocalResourceId(),
				ReporterVersion:    resourceEvent.ReporterVersion(), //nolint:staticcheck
			},
			ResourceData: resourceEvent.Data(),
		},
	}, nil
}

func makeEventType(eventType, resourceType, operation string) string {
	return fmt.Sprintf("redhat.inventory.%s.%s.%s", eventType, resourceType, operation)
}

func makeEventSubject(eventType, resourceType, resourceId string) string {
	return "/" + strings.Join([]string{eventType, resourceType, resourceId}, "/")
}
