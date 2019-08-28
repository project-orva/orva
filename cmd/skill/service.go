package main

import (
	"context"

	grpcSkill"github.com/GuyARoss/project-orva/pkg/grpc-skill"
)

// CreateSkill creates a new skill
func (server *ServiceRequest) CreateSkill(
	ctx context.Context,
	*grpcSkill.Skill,
) (*grpcSkill.SkillResponse, error) {
	return nil, nil
}

// Delete Skill deletes a skill
func (server *ServiceRequest) DeleteSkill(ctx context.Context, *grpcSkill.Skill) (*grpcSkill.SkillResponse, error) {
	return nil, nil
}

// UpdateSkill updates a skill 
func (server *ServiceRequest) UpdateSkill(ctx context.Context, *grpcSkill.Skill) (*grpcSkill.SkillResponse, error) {
	return nil, nil
}

// QuerySkill queries a skill given the id
func (server *ServiceRequest) QuerySkill(ctx context.Context, *grpcSkill.QueryRequest) (*grpcSkill.Skill, error) {
	return nil, nil
}