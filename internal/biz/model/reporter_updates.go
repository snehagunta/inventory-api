package model

import (
	"errors"
)

var (
	ErrAtLeastOneVersionRequired = errors.New("at least one version (common or reporter) must be provided")
)

type ResourceVersions struct {
	commonVersion   *Version
	reporterVersion *Version
}

func NewResourceVersions(commonVersion *Version, reporterVersion *Version) (ResourceVersions, error) {
	if commonVersion == nil && reporterVersion == nil {
		return ResourceVersions{}, ErrAtLeastOneVersionRequired
	}

	return ResourceVersions{
		commonVersion:   commonVersion,
		reporterVersion: reporterVersion,
	}, nil
}

func (rv ResourceVersions) CommonVersion() *Version {
	return rv.commonVersion
}

func (rv ResourceVersions) ReporterVersion() *Version {
	return rv.reporterVersion
}

func (rv ResourceVersions) HasCommonVersion() bool {
	return rv.commonVersion != nil
}

func (rv ResourceVersions) HasReporterVersion() bool {
	return rv.reporterVersion != nil
}

type ReporterUpdates struct {
	key              ReporterResourceKey
	resourceVersions ResourceVersions
}

func NewReporterUpdates(key ReporterResourceKey, resourceVersions ResourceVersions) (ReporterUpdates, error) {
	return ReporterUpdates{
		key:              key,
		resourceVersions: resourceVersions,
	}, nil
}

func (ru ReporterUpdates) Key() ReporterResourceKey {
	return ru.key
}

func (ru ReporterUpdates) ResourceVersions() ResourceVersions {
	return ru.resourceVersions
}
