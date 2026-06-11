package model

import (
	"strings"
)

// RelationsTuple represents a stored relationship fact.
// Structurally identical to Relationship but semantically distinct:
// a Relationship is a query ("does this hold?"), a RelationsTuple is a persisted fact.
type RelationsTuple struct {
	object   ResourceReference
	relation Relation
	subject  SubjectReference
}

func NewRelationsTuple(object ResourceReference, relation Relation, subject SubjectReference) RelationsTuple {
	return RelationsTuple{
		object:   object,
		relation: relation,
		subject:  subject,
	}
}

func (rt RelationsTuple) Object() ResourceReference { return rt.object }
func (rt RelationsTuple) Relation() Relation        { return rt.relation }
func (rt RelationsTuple) Subject() SubjectReference { return rt.subject }

const (
	WorkspaceRelation        = "workspace"
	AllowedWorkspacesRelation = "allowed_workspaces"
	BillingAccountRelation   = "billing_account"
	ParentRelation           = "parent"
	RbacNamespace            = "rbac"
	FeatureNamespace         = "features"
)

// NewAllowedWorkspacesTuple creates a tuple: feature/service:{id}#allowed_workspaces@rbac/workspace:{wsID}
func NewAllowedWorkspacesTuple(workspaceID string, key ReporterResourceKey) RelationsTuple {
	reporter := NewReporterReference(key.ReporterType(), nil)
	object := NewResourceReference(
		key.ResourceType(),
		key.LocalResourceId(),
		&reporter,
	)

	wsSubjectId := DeserializeLocalResourceId(workspaceID)
	wsReporterType := DeserializeReporterType(RbacNamespace)
	wsReporter := NewReporterReference(wsReporterType, nil)
	wsResource := NewResourceReference(
		DeserializeResourceType(WorkspaceRelation),
		wsSubjectId,
		&wsReporter,
	)
	subject := NewSubjectReferenceWithoutRelation(wsResource)

	return RelationsTuple{
		object:   object,
		relation: DeserializeRelation(AllowedWorkspacesRelation),
		subject:  subject,
	}
}

// NewBillingAccountTuple creates a tuple: feature/service:{id}#billing_account@feature/billing_account:{baID}
func NewBillingAccountTuple(billingAccountID string, key ReporterResourceKey) RelationsTuple {
	reporter := NewReporterReference(key.ReporterType(), nil)
	object := NewResourceReference(
		key.ResourceType(),
		key.LocalResourceId(),
		&reporter,
	)

	baSubjectId := DeserializeLocalResourceId(billingAccountID)
	baReporterType := DeserializeReporterType(FeatureNamespace)
	baReporter := NewReporterReference(baReporterType, nil)
	baResource := NewResourceReference(
		DeserializeResourceType(BillingAccountRelation),
		baSubjectId,
		&baReporter,
	)
	subject := NewSubjectReferenceWithoutRelation(baResource)

	return RelationsTuple{
		object:   object,
		relation: DeserializeRelation(BillingAccountRelation),
		subject:  subject,
	}
}

// NewParentServiceTuple creates a tuple: feature/service:{id}#parent@feature/service:{parentID}
func NewParentServiceTuple(parentServiceID string, key ReporterResourceKey) RelationsTuple {
	reporter := NewReporterReference(key.ReporterType(), nil)
	object := NewResourceReference(
		key.ResourceType(),
		key.LocalResourceId(),
		&reporter,
	)

	parentSubjectId := DeserializeLocalResourceId(parentServiceID)
	parentReporterType := DeserializeReporterType(FeatureNamespace)
	parentReporter := NewReporterReference(parentReporterType, nil)
	parentResource := NewResourceReference(
		DeserializeResourceType("service"),
		parentSubjectId,
		&parentReporter,
	)
	subject := NewSubjectReferenceWithoutRelation(parentResource)

	return RelationsTuple{
		object:   object,
		relation: DeserializeRelation(ParentRelation),
		subject:  subject,
	}
}

func NewWorkspaceRelationsTuple(workspaceID string, key ReporterResourceKey) RelationsTuple {
	reporter := NewReporterReference(key.ReporterType(), nil)
	object := NewResourceReference(
		key.ResourceType(),
		key.LocalResourceId(),
		&reporter,
	)

	workspaceSubjectId := DeserializeLocalResourceId(workspaceID)
	workspaceReporterType := DeserializeReporterType(RbacNamespace)
	workspaceReporter := NewReporterReference(workspaceReporterType, nil)
	workspaceResource := NewResourceReference(
		DeserializeResourceType(WorkspaceRelation),
		workspaceSubjectId,
		&workspaceReporter,
	)
	subject := NewSubjectReferenceWithoutRelation(workspaceResource)

	return RelationsTuple{
		object:   object,
		relation: DeserializeRelation(strings.ToLower(WorkspaceRelation)),
		subject:  subject,
	}
}
