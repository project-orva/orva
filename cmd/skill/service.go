package main

import (
	"context"

	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	pgSkill "github.com/GuyARoss/project-orva/pkg/pg-schemas/skill"
	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

// ServiceRequest cmd service request required by grpc
type ServiceRequest struct {
	Creds *pgdb.DbCreds
}

// CreateSkill creates a new skill.
func (server *ServiceRequest) CreateSkill(
	ctx context.Context,
	skill *grpcSkill.Skill,
) (
	*grpcSkill.SuccessResponse,
	error,
) {
	pgReq := &pgSkill.Request{server.Creds}

	skillReq := &pgSkill.Skill{
		ID: skill.ID,
		Endpoint: skill.Endpoint,
		AccessLevel: skill.AccessLevel,
	}

	err := pgReq.Create(skillReq)
	if err != nil {
		return &grpcSkill.SuccessResponse{
			Success: false,
		}, err
	}

	return &grpcSkill.SuccessResponse{
		Success: true,
	}, nil
}

// DeleteSkill deletes a skill from the skill db.
func (server *ServiceRequest) DeleteSkill(
	ctx context.Context,
	skill *grpcSkill.Skill,
) (
	*grpcSkill.SuccessResponse,
	error,
) {
	pgReq := &pgSkill.Request{server.Creds}

	skillReq := &pgSkill.Skill{
		ID: skill.ID,
		Endpoint: skill.Endpoint,
		AccessLevel: skill.AccessLevel,
	}

	err := pgReq.Delete(skillReq)
	if err != nil {
		return &grpcSkill.SuccessResponse{
			Success: false,
		}, err
	}

	return &grpcSkill.SuccessResponse{
		Success: true,
	}, nil
}

// UpdateSkill updates a skill.
func (server *ServiceRequest) UpdateSkill(
	ctx context.Context,
	skill *grpcSkill.Skill,
) (
	*grpcSkill.SuccessResponse,
	error,
) {
	pgReq := &pgSkill.Request{server.Creds}
	skillReq := &pgSkill.Skill{
		ID: skill.ID,
		Endpoint: skill.Endpoint,
		AccessLevel: skill.AccessLevel,
	}

	err := pgReq.Update(skillReq)
	if err != nil {
		return &grpcSkill.SuccessResponse{
			Success: false,
		}, err
	}

	return &grpcSkill.SuccessResponse{
		Success: true,
	}, nil
}

// QuerySkill queries a skill given the id.
func (server *ServiceRequest) QuerySkill(
	ctx context.Context,
	skill *grpcSkill.QueryRequest,
) (
	*grpcSkill.Skill,
	error,
) {
	pgReq := &pgSkill.Request{server.Creds}

	pgSkill, err := pgReq.FindByID(skill.ID)
	if err != nil {
		return nil, err
	}
	
	return &grpcSkill.Skill{
		ID: pgSkill.ID,
		Endpoint: pgSkill.Endpoint,
		AccessLevel: pgSkill.AccessLevel,
	}, nil
}
