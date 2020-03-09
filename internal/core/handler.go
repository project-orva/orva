package core

import (
	grpcAccount "github.com/GuyARoss/project-orva/pkg/grpc/account"
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	pgdb "github.com/GuyARoss/project-orva/pkg/pgdb"
	"github.com/GuyARoss/project-orva/pkg/orva"
)

type RoutineRequest struct {
	AccountClient grpcAccount.GrpcAccountClient
	SpeechClient  grpcSpeech.GrpcSpeechClient
	SkillClient   grpcSkill.GrpcSkillClient
	PgCreds			*pgdb.DbCreds	
}

func initRequests(accountAddress string, speechAddress string, skillAddress string) *RoutineRequest {
	return &RoutineRequest{
		AccountClient: grpcAccount.CreateClientConn(accountAddress),
		SpeechClient:  grpcSpeech.CreateClientConn(speechAddress),
		SkillClient:   grpcSkill.CreateClientConn(skillAddress),
	}
}

func invokeRoutineHandlers(ctx *orva.SessionContext) {
	// todo: make these go routines.

	req.AccountRoutineHandler(ctx)
	req.SkillRoutineHandler(ctx)
	req.SpeechRoutineHandler(ctx)
}

// CoreHandler unifies orva routines
func (req *RoutineRequest) CoreHandler(ctx *orva.SessionContext) {
	invokeRoutineHandlers(ctx)

	// todo: save routine ctx to long-term memory.
}
