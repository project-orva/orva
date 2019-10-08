package main

import (
	"context"

	grpcAccount "github.com/GuyARoss/project-orva/pkg/grpc/account"
	pgAccount "github.com/GuyARoss/project-orva/pkg/pg-schemas/account"
	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

type ServiceRequest struct {
	Creds *pgdb.DbCreds
}

// RetrieveFromId gets a profile based on the provided ID
func (request *ServiceRequest) RetrieveFromId(ctx context.Context, accReq *grpcAccount.AccountRequest) (*grpcAccount.Account, error) {
	pgReq := &pgAccount.Request{request.Creds}
	resp, err := pgReq.FindByID(accReq.ID)
	if err != nil {
		return &grpcAccount.Account{
			Type:        resp.Type,
			AccessLevel: resp.AccessLevel,
			Name:        resp.Name,
			ID:          resp.ID,
		}, nil
	}

	return nil, err
}

// UpdateAccount updates a profile
func (request *ServiceRequest) UpdateAccount(ctx context.Context, acc *grpcAccount.Account) (*grpcAccount.SuccessResponse, error) {
	return nil, nil
}

// CreateAccount creates a new profile
func (request *ServiceRequest) CreateAccount(ctx context.Context, acc *grpcAccount.Account) (*grpcAccount.SuccessResponse, error) {
	return nil, nil
}
