package resources

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pbv2 "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta2/resources"
	"github.com/project-kessel/inventory-api/internal/biz/resources"
	"google.golang.org/protobuf/types/known/structpb"
)

type ResourceService struct {
	pbv2.UnimplementedKesselResourceServiceServer
	Ctl *resources.Usecase
}

func New(c *resources.Usecase) *ResourceService {
	return &ResourceService{
		Ctl: c,
	}
}

func (c *ResourceService) ReportResource(ctx context.Context, r *pbv2.ReportResourceRequest) (*pbv2.ReportResourceResponse, error) {
	log.Info("Processing ReportResource request", r)

	// Validate and parse resource data
	if r.Resource == nil {
		return nil, fmt.Errorf("resource field is required")
	}

	reporterData, err := createReporterData(r.Resource.ReporterData)
	if err != nil {
		return nil, fmt.Errorf("failed to process reporter data: %w", err)
	}

	var commonResourceData *structpb.Struct
	if r.Resource.CommonResourceData != nil {
		commonResourceData = r.Resource.CommonResourceData
	}

	processedResource := &pbv2.Resource{
		ResourceType:       r.Resource.ResourceType,
		ReporterData:       reporterData,
		CommonResourceData: commonResourceData,
	}

	// Simulate resource storage operation
	log.Info("Successfully processed resource", processedResource)

	return &pbv2.ReportResourceResponse{}, nil
}

func (c *ResourceService) DeleteResource(ctx context.Context, r *pbv2.DeleteResourceRequest) (*pbv2.DeleteResourceResponse, error) {
	log.Info("Processing DeleteResource request", r)

	if r.LocalResourceId == "" || r.ReporterType == "" {
		return nil, fmt.Errorf("local_resource_id and reporter_type are required")
	}

	// Simulate deletion operation
	log.Info("Successfully deleted resource", r.LocalResourceId)

	return &pbv2.DeleteResourceResponse{}, nil
}

func createReporterData(reporter *pbv2.ReporterData) (*pbv2.ReporterData, error) {
	if reporter == nil {
		return nil, fmt.Errorf("reporter data is missing")
	}

	return &pbv2.ReporterData{
		ReporterType:       reporter.ReporterType,
		ReporterInstanceId: reporter.ReporterInstanceId,
		ReporterVersion:    reporter.ReporterVersion,
		LocalResourceId:    reporter.LocalResourceId,
		ApiHref:            reporter.ApiHref,
		ConsoleHref:        reporter.ConsoleHref,
		ResourceData:       reporter.ResourceData,
	}, nil
}
