package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/project-kessel/inventory-api/internal"
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
	Operation     internal.EventOperationType `gorm:"type:varchar(255);not null"`
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
