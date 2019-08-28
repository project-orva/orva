package main

import (
	"context"

	grpcMemory "github.com/GuyARoss/project-orva/pkg/grpc/memory"
)

// GetLatestFromUser gets the latest memory from the desired user
func (server *ServiceRequest) GetLatestFromUser(ctx context.Context, memoryRequest *grpcMemory.MemoryRequest) (*grpcMemory.MemorySet, error) {
	return nil, nil
}

// GetLatestFromDevice gets the latest memory from the desired device
func (server *ServiceRequest) GetLatestFromDevice(ctx context.Context, memoryRequest *grpcMemory.MemoryRequest) (*grpcMemory.MemorySet, error) {
	return nil, nil
}

// CreateMemory creates a new memory instance
func (server *ServiceRequest) CreateMemory(ctx context.Context, memoryRequest *grpcMemory.MemorySet) (*grpcMemory.MemoryRequest, error) {
	return nil, nil
}
