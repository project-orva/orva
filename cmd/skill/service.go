package main

import (
	"context"

	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
)

// CreateSkill creates a new skill.
func (server *ServiceRequest) CreateSkill(ctx context.Context, skill *grpcSkill.Skill) (*grpcSkill.SuccessResponse, error) {
	return nil, nil
}

// DeleteSkill deletes a skill from the skill db.
func (server *ServiceRequest) DeleteSkill(ctx context.Context, skill *grpcSkill.Skill) (*grpcSkill.SuccessResponse, error) {
	return nil, nil
}

// UpdateSkill updates a skill.
func (server *ServiceRequest) UpdateSkill(ctx context.Context, skill *grpcSkill.Skill) (*grpcSkill.SuccessResponse, error) {
	return nil, nil
}

// QuerySkill queries a skill given the id.
func (server *ServiceRequest) QuerySkill(ctx context.Context, skill *grpcSkill.QueryRequest) (*grpcSkill.Skill, error) {
	return nil, nil
}
