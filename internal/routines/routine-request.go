package routines

import (
	grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
)

// RoutineRequest base request settings for routines
type RoutineRequest struct {
	ProfileClient grpcProfile.GrpcProfileClient
	SpeechClient  grpcSpeech.GrpcSpeechClient
	SkillClient   grpcSkill.GrpcSkillClient
}

func initRequests(profileAddress string, speechAddress string, skillAddress string) *RoutineRequest {
	return &RoutineRequest{
		ProfileClient: grpcProfile.CreateClientConn(profileAddress),
		SpeechClient:  grpcSpeech.CreateClientConn(speechAddress),
		SkillClient:   grpcSkill.CreateClientConn(skillAddress),
	}
}
