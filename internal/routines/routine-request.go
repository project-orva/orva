package routines

import (
	grpcAccount "github.com/GuyARoss/project-orva/pkg/grpc/account"
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
)

// RoutineRequest base request settings for routines
type RoutineRequest struct {
	AccountClient grpcAccount.GrpcAccountClient
	SpeechClient  grpcSpeech.GrpcSpeechClient
	SkillClient   grpcSkill.GrpcSkillClient
}

func initRequests(accountAddress string, speechAddress string, skillAddress string) *RoutineRequest {
	return &RoutineRequest{
		AccountClient: grpcAccount.CreateClientConn(accountAddress),
		SpeechClient:  grpcSpeech.CreateClientConn(speechAddress),
		SkillClient:   grpcSkill.CreateClientConn(skillAddress),
	}
}
