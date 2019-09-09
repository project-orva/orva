package main

import (
	"context"

	grpcProfile "github.com/GuyARoss/project-orva/grpc/profile"
)

// FindProfileByAccountID finds a profile given the account ID.
func (req *ServiceRequest) FindProfileByAccountID(ctx context.Context, profileReq) (*grpcProfile.Profile, error) {

}

// CreateProfile creates a new profile.
func (req *ServiceRequest) CreateProfile(ctx context.Context, profile *grpcProfile.Profile) (*grpcAccount, error) {

} 