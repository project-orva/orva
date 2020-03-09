package main

import (
	"context"

	grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"
)

type ServiceRequest struct {
}

// FindProfileByUserID finds a profile given the account ID.
func (req *ServiceRequest) FindProfileByUserID(ctx context.Context, profileReq *grpcProfile.ProfileRequest) (*grpcProfile.Profile, error) {
	return nil, nil
}

// CreateProfile creates a new profile.
func (req *ServiceRequest) CreateProfile(ctx context.Context, profile *grpcProfile.Profile) (*grpcProfile.CreationResponse, error) {
	return nil, nil
}
