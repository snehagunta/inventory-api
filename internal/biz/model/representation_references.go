package model

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// RepresentationReference is an immutable value object representing a reference between resources and their representations.
// It follows DDD principles where value objects are immutable and should be created
// through factory methods that enforce validation rules.
// Note: Fields are exported for GORM compatibility but should not be modified directly.
type RepresentationReference struct {
	ResourceID uuid.UUID `gorm:"type:uuid;column:resource_id;index:rep_ref_unique_idx,unique;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	LocalResourceID       string `gorm:"size:128;column:local_resource_id;index:rep_ref_unique_idx,unique"`
	ReporterType          string `gorm:"size:128;column:reporter_type;index:rep_ref_unique_idx,unique;not null"`
	ResourceType          string `gorm:"size:128;column:resource_type;index:rep_ref_unique_idx,unique;not null"`
	ReporterInstanceID    string `gorm:"size:128;column:reporter_instance_id;index:rep_ref_unique_idx,unique;not null"`
	RepresentationVersion int    `gorm:"type:bigint;column:representation_version;index:rep_ref_unique_idx,unique;check:representation_version >= 0"`
	Generation            int    `gorm:"type:bigint;column:generation;index:rep_ref_unique_idx,unique;check:generation >= 0"`
	Tombstone             bool   `gorm:"column:tombstone"`
}

// NewRepresentationReference creates a RepresentationReference
// This enforces immutability by validating all inputs and creating a valid instance
func NewRepresentationReference(
	resourceID uuid.UUID,
	localResourceID string,
	reporterType string,
	resourceType string,
	reporterInstanceID string,
	representationVersion int,
	generation int,
	tombstone bool,
) (*RepresentationReference, error) {
	rr := &RepresentationReference{
		ResourceID:            resourceID,
		LocalResourceID:       localResourceID,
		ReporterType:          reporterType,
		ResourceType:          resourceType,
		ReporterInstanceID:    reporterInstanceID,
		RepresentationVersion: representationVersion,
		Generation:            generation,
		Tombstone:             tombstone,
	}

	// Validate the instance
	if err := validateRepresentationReference(rr); err != nil {
		return nil, err
	}

	return rr, nil
}

// validateRepresentationReference validates a RepresentationReference instance
// This function is used internally by factory methods to ensure consistency
func validateRepresentationReference(rr *RepresentationReference) error {
	if rr.ResourceID == uuid.Nil {
		return ValidationError{Field: "ResourceID", Message: "cannot be empty"}
	}
	if rr.LocalResourceID == "" || strings.TrimSpace(rr.LocalResourceID) == "" {
		return ValidationError{Field: "LocalResourceID", Message: "cannot be empty"}
	}
	if len(rr.LocalResourceID) > MaxLocalResourceIDLength {
		return ValidationError{Field: "LocalResourceID", Message: fmt.Sprintf("exceeds maximum length of %d characters", MaxLocalResourceIDLength)}
	}
	if rr.ReporterType == "" || strings.TrimSpace(rr.ReporterType) == "" {
		return ValidationError{Field: "ReporterType", Message: "cannot be empty"}
	}
	if len(rr.ReporterType) > MaxReporterTypeLength {
		return ValidationError{Field: "ReporterType", Message: fmt.Sprintf("exceeds maximum length of %d characters", MaxReporterTypeLength)}
	}
	if rr.ResourceType == "" || strings.TrimSpace(rr.ResourceType) == "" {
		return ValidationError{Field: "ResourceType", Message: "cannot be empty"}
	}
	if len(rr.ResourceType) > MaxResourceTypeLength {
		return ValidationError{Field: "ResourceType", Message: fmt.Sprintf("exceeds maximum length of %d characters", MaxResourceTypeLength)}
	}
	if rr.ReporterInstanceID == "" || strings.TrimSpace(rr.ReporterInstanceID) == "" {
		return ValidationError{Field: "ReporterInstanceID", Message: "cannot be empty"}
	}
	if len(rr.ReporterInstanceID) > MaxReporterInstanceIDLength {
		return ValidationError{Field: "ReporterInstanceID", Message: fmt.Sprintf("exceeds maximum length of %d characters", MaxReporterInstanceIDLength)}
	}
	if rr.RepresentationVersion < MinRepresentationVersion {
		return ValidationError{Field: "RepresentationVersion", Message: fmt.Sprintf("must be >= %d", MinRepresentationVersion)}
	}
	if rr.Generation < MinGenerationValue {
		return ValidationError{Field: "Generation", Message: fmt.Sprintf("must be >= %d", MinGenerationValue)}
	}
	return nil
}
