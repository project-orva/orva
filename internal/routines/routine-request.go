package routines

import (
	grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
) 

// RoutineRequest base request settings for routines 
type RoutineRequest struct {
	ProfileClient grpcProfile.GrpcProfileClient
	SpeechClient grpcSpeech.GrpcSpeechClient
	SkillClient grpcSkill.GrpcSkillClient
}
