package main

import (
	"context"

	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
)

type ServiceRequest struct {
}

// ProcessStatement processes the inputted statement, outputs a response.
func (server *ServiceRequest) ProcessStatement(ctx context.Context, req *grpcCore.Request) (*grpcCore.Response, error) {
	return nil, nil
}
