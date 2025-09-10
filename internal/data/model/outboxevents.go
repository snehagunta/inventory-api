package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/project-kessel/inventory-api/internal"
	bizmodel "github.com/project-kessel/inventory-api/internal/biz/model"
	"gorm.io/gorm"
)

type AggregateType string

const (
	TupleAggregateType    AggregateType = "kessel.tuples"
	ResourceAggregateType AggregateType = "kessel.resources"
)

type OutboxEvent struct {
	ID            uuid.UUID                   `gorm:"type:uuid;primarykey;not null"`
	AggregateType AggregateType               `gorm:"column:aggregatetype;type:varchar(255);not null"`
	AggregateID   string                      `gorm:"column:aggregateid;type:varchar(255);not null"`
	Operation     bizmodel.EventOperationType `gorm:"type:varchar(255);not null"`
	TxId          string                      `gorm:"column:txid;type:varchar(255)"`
	Payload       internal.JsonObject
}

func (r *OutboxEvent) BeforeCreate(db *gorm.DB) error {
	var err error
	if r.ID == uuid.Nil {
		r.ID, err = uuid.NewV7()
		if err != nil {
			return fmt.Errorf("failed to generate uuid: %w", err)
		}
	}
	return nil
}

type ResourceEvent struct {
	Specversion     string      `json:"specversion"`
	Type            string      `json:"type"`
	Source          string      `json:"source"`
	Id              string      `json:"id"`
	Subject         string      `json:"subject"`
	Time            time.Time   `json:"time"`
	DataContentType string      `json:"datacontenttype"`
	Data            interface{} `json:"data"`
}

type EventResourceData struct {
	Metadata     EventResourceMetadata `json:"metadata"`
	ReporterData EventResourceReporter `json:"reporter_data"`
	ResourceData internal.JsonObject   `json:"resource_data,omitempty"`
}

type EventRelationshipData struct {
	Metadata     EventRelationshipMetadata `json:"metadata"`
	ReporterData EventRelationshipReporter `json:"reporter_data"`
	ResourceData internal.JsonObject       `json:"resource_data,omitempty"`
}

type EventResourceMetadata struct {
	Id           string               `json:"id"`
	ResourceType string               `json:"resource_type"`
	OrgId        string               `json:"org_id"`
	CreatedAt    *time.Time           `json:"created_at,omitempty"`
	UpdatedAt    *time.Time           `json:"updated_at,omitempty"`
	DeletedAt    *time.Time           `json:"deleted_at,omitempty"`
	WorkspaceId  string               `json:"workspace_id"`
	Labels       []EventResourceLabel `json:"labels,omitempty"`
}

type EventResourceLabel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EventResourceReporter struct {
	ReporterInstanceId string  `json:"reporter_instance_id"`
	ReporterType       string  `json:"reporter_type"`
	ConsoleHref        string  `json:"console_href"`
	ApiHref            string  `json:"api_href"`
	LocalResourceId    string  `json:"local_resource_id"`
	ReporterVersion    *string `json:"reporter_version"`
}

type EventRelationshipMetadata struct {
	Id               string     `json:"id"`
	RelationshipType string     `json:"relationship_type"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type EventRelationshipReporter struct {
	ReporterType           string `json:"reporter_type"`
	SubjectLocalResourceId string `json:"subject_local_resource_id"`
	ObjectLocalResourceId  string `json:"object_local_resource_id"`
	ReporterVersion        string `json:"reporter_version"`
	ReporterInstanceId     string `json:"reporter_instance_id"`
}

func NewOutboxEventsFromResourceEvent(domainResourceEvent bizmodel.ResourceReportEvent, operationType bizmodel.EventOperationType, txid string) (*OutboxEvent, *OutboxEvent, error) {
	var tuplePayload internal.JsonObject
	var tupleEvent *OutboxEvent

	// Build resource event
	payload, err := convertResourceToResourceEvent(domainResourceEvent, operationType)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert resource to payload: %w", err)
	}

	resourceEvent := &OutboxEvent{
		Operation:     operationType,
		AggregateType: ResourceAggregateType,
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

	tupleEvent = &OutboxEvent{
		Operation:     operationType,
		AggregateType: TupleAggregateType,
		AggregateID:   domainResourceEvent.Id().String(),
		TxId:          txid,
		Payload:       tuplePayload,
	}

	return resourceEvent, tupleEvent, nil
}

func convertResourceToResourceEvent(resourceReportEvent bizmodel.ResourceReportEvent, operationType bizmodel.EventOperationType) (internal.JsonObject, error) {
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

func newResourceEvent(operationType bizmodel.EventOperationType, resourceEvent *bizmodel.ResourceReportEvent) (*ResourceEvent, error) {
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
	case bizmodel.OperationTypeCreated:
		createdAt = resourceEvent.CreatedAt()
		reportedTime = *createdAt
	case bizmodel.OperationTypeUpdated:
		updatedAt = resourceEvent.UpdatedAt()
		reportedTime = *updatedAt
	case bizmodel.OperationTypeDeleted:
		deletedAt = &now
		reportedTime = *deletedAt
	}

	return &ResourceEvent{
		Specversion:     "1.0",
		Type:            MakeEventType(eventType, resourceEvent.ResourceType(), string(operationType.OperationType())),
		Source:          "", // TODO: inventory uri
		Id:              eventId.String(),
		Subject:         MakeEventSubject(eventType, resourceEvent.ResourceType(), resourceEvent.Id().String()),
		Time:            reportedTime,
		DataContentType: "application/json",
		Data: EventResourceData{
			Metadata: EventResourceMetadata{
				Id:           resourceEvent.Id().String(),
				ResourceType: resourceEvent.ResourceType(),
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
				DeletedAt:    deletedAt,
				WorkspaceId:  resourceEvent.WorkspaceId(),
			},
			ReporterData: EventResourceReporter{
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

func MakeEventType(eventType, resourceType, operation string) string {
	return fmt.Sprintf("redhat.inventory.%s.%s.%s", eventType, resourceType, operation)
}

func MakeEventSubject(eventType, resourceType, resourceId string) string {
	return "/" + strings.Join([]string{eventType, resourceType, resourceId}, "/")
}
