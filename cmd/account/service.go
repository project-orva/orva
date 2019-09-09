package main

import (
	"context"

	grpcAccount "github.com/GuyARoss/project-orva/pkg/grpc/account"
)

// RetrieveFromId gets a profile based on the provided ID
func (request *ServiceRequest) RetrieveFromId(ctx context.Context, accReq *grpcAccount.AccountRequest) (*grpcAccount.Account, error) {
	return nil, nil
}

// UpdateAccount updates a profile
func (request *ServiceRequest) UpdateAccount(ctx context.Context, acc *grpcAccount.Account) (*grpcAccount.SuccessResponse, error) {
	return nil, nil
}

// CreateAccount creates a new profile
func (request *ServiceRequest) CreateAccount(ctx context.Context, acc *grpcAccount.Account) (*grpcAccount.SuccessResponse, error) {
	return nil, nil
}
