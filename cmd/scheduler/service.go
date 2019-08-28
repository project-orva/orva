package main

import (
	"context"

	grpcScheduler "github.com/GuyARoss/project-orva/pkg/grpc-scheduler"
)

// CreateJob creates a new http web job.
func (server *ServiceRequest) CreateJob(ctx context.Context, job *grpcScheduler.Job) (*grpcScheduler.SuccessResponse, error) {
	return nil, nil
}

// QueryJob queries existing jobs by ID.
func (server *ServiceRequest) QueryJob(ctx context.Context, query *grpcScheduler.JobQuery) (*grpcScheduler.JobQueryResponse, error) {
	return nil, nil
}
