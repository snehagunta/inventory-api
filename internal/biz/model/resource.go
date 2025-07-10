package model

import "github.com/google/uuid"

type Resource struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type                     string    `gorm:"size:128"`
	ConsistencyToken         string
	RepresentationReferences RepresentationReference
}

func (Resource) TableName() string {
	return "resource"
}
