package main

import (
	"context"

	grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"
)

// RetrieveFromId gets a profile based on the provided ID
func (request *ServiceRequest) RetrieveFromId(ctx context.Context, profileRequest *grpcProfile.ProfileRequest) (*grpcProfile.Profile, error) {
	return nil, nil
}

// UpdateAccount updates a profile
func (request *ServiceRequest) UpdateAccount(ctx context.Context, profile *grpcProfile.Profile) (*grpcProfile.SuccessResponse, error) {
	return nil, nil
}

// CreateAccount creates a new profile
func (request *ServiceRequest) CreateAccount(ctx context.Context, profile *grpcProfile.Profile) (*grpcProfile.SuccessResponse, error) {
	return nil, nil
}
