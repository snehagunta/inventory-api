package model

// Database field size constants
// These constants define the maximum lengths for various database fields
// to ensure consistency between GORM struct tags and validation logic.
const (
	// Standard field sizes
	MaxFieldSize128 = 128 // For most string fields like IDs, types, etc.
	MaxFieldSize512 = 512 // For URL fields like APIHref, ConsoleHref

	// Specific field size constants for better readability
	MaxLocalResourceIDLength    = MaxFieldSize128
	MaxReporterTypeLength       = MaxFieldSize128
	MaxResourceTypeLength       = MaxFieldSize128
	MaxReporterInstanceIDLength = MaxFieldSize128
	MaxReporterVersionLength    = MaxFieldSize128
	MaxAPIHrefLength            = MaxFieldSize512
	MaxConsoleHrefLength        = MaxFieldSize512

	// Minimum values for validation
	MinVersionValue          = 0 // Version can be zero or positive (>= 0)
	MinGenerationValue       = 0 // Generation can be zero or positive (>= 0)
	MinCommonVersion         = 0 // CommonVersion can be zero or positive (>= 0)
	MinRepresentationVersion = 0 // RepresentationVersion can be zero or positive (>= 0)
)

// Column name constants
const (
	// CommonRepresentation columns
	ColumnResourceID                 = "id"
	ColumnResourceType               = "resource_type"
	ColumnVersion                    = "version"
	ColumnReportedByReporterType     = "reported_by_reporter_type"
	ColumnReportedByReporterInstance = "reported_by_reporter_instance"
	ColumnData                       = "data"

	// ReporterRepresentation columns
	ColumnLocalResourceID    = "local_resource_id"
	ColumnReporterType       = "reporter_type"
	ColumnReporterInstanceID = "reporter_instance_id"
	ColumnGeneration         = "generation"
	ColumnAPIHref            = "api_href"
	ColumnConsoleHref        = "console_href"
	ColumnCommonVersion      = "common_version"
	ColumnTombstone          = "tombstone"
	ColumnReporterVersion    = "reporter_version"

	// RepresentationReference columns
	ColumnRepresentationReferenceResourceID = "resource_id"
	ColumnRepresentationVersion             = "representation_version"
)

// Index names
const (
	ReporterRepresentationUniqueIndex  = "reporter_rep_unique_idx"
	RepresentationReferenceUniqueIndex = "unique_rep_ref_idx"
)

// Database type constants
const (
	DBTypeText   = "text"
	DBTypeBigInt = "bigint"
	DBTypeJSONB  = "jsonb"
)
