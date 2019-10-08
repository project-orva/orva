package main

import (
	"context"

	grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"
	pgProfile "github.com/GuyARoss/project-orva/pkg/pg-schemas/profile"
	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

// ServiceRequest request struct for remote procedures.
type ServiceRequest struct {
	Creds *pgdb.DbCreds
}

// FindProfileByAccountID finds a profile given the account ID.
func (req *ServiceRequest) FindProfileByAccountID(ctx context.Context, profileReq *grpcProfile.ProfileRequest) (*grpcProfile.Profile, error) {
	pgReq := &pgProfile.Request{req.Creds}
	p, err := pgReq.FindByID(profileReq.ID)
	if err != nil {
		return nil, err
	}

	return &grpcProfile.Profile{
		BlockAddress: p.BlockAddress,
		AccountID:    p.AccountID,
		CreatedOn:    p.CreatedOn,
	}, nil
}

// CreateProfile creates a new profile.
func (req *ServiceRequest) CreateProfile(ctx context.Context, profile *grpcProfile.Profile) (*grpcProfile.CreationResponse, error) {
	pgReq := &pgProfile.Request{req.Creds}
	err := pgReq.Create(&pgProfile.Profile{
		AccountID:    profile.AccountID,
		BlockAddress: profile.BlockAddress,
		CreatedOn:    profile.CreatedOn,
	})

	return &grpcProfile.CreationResponse{
		Success: err == nil,
	}, err
}
