package main

import (
	"context"

	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	"github.com/GuyARoss/project-orva/pkg/orva"
	pgDevice "github.com/GuyARoss/project-orva/pkg/pg-schemas/device"
	pgUser "github.com/GuyARoss/project-orva/pkg/pg-schemas/user"
	"github.com/GuyARoss/project-orva/pkg/pgdb"
	"github.com/GuyARoss/project-orva/pkg/utilities/mappings"
)

type RoutineRequest struct {
	SpeechClient grpcSpeech.GrpcSpeechClient
	PgCreds      *pgdb.DbCreds
}

// CoreHandler unifies orva routines
func (req *RoutineRequest) CoreHandler(ctx *orva.SessionContext) {
	req.invokeRoutineHandlers(ctx)

	// todo: save routine ctx to long-term memory.
}

func (req *RoutineRequest) invokeRoutineHandlers(ctx *orva.SessionContext) {
	// todo: make these go routines.

	req.AccountRoutineHandler(ctx)
	req.SkillRoutineHandler(ctx)
	req.SpeechRoutineHandler(ctx)
}

// AccountRoutineHandler handles the profiles routine
func (req *RoutineRequest) AccountRoutineHandler(ctx *orva.SessionContext) {
	user, userErr := req.verifyUser(ctx.InitialInput.UserID)
	device, deviceErr := req.verifyDevice(ctx.InitialInput.DeviceID)

	if userErr != nil {
		ctx.UserAccessLvl = orva.AnonAccess
		resp := &orva.Response{
			Statement: "Could not validate your user access",
		}

		ctx.Append(resp)
	} else {
		ctx.UserAccessLvl = orva.AccessType(user.AccessLevel)
	}

	if deviceErr != nil {
		ctx.DeviceAccessLvl = orva.AnonAccess

		resp := &orva.Response{
			Statement: "Could not validate your device access",
		}

		ctx.Append(resp)
	} else {
		ctx.DeviceAccessLvl = orva.AccessType(device.AccessLevel)
	}
}

// account_checkUser checks user profile account under the account routine handler.
func (req *RoutineRequest) verifyUser(UID string) (*pgUser.User, error) {
	pgReq := &pgUser.Request{req.PgCreds}

	return pgReq.FindByID(UID)
}

// account_checkDevice checks device account under the account routine handler.
func (req *RoutineRequest) verifyDevice(DID string) (*pgDevice.Device, error) {
	pgReq := &pgDevice.Request{req.PgCreds}

	return pgReq.FindByID(DID)
}

func determineSkill(initialStatement string) (string, error) {
	return "http://localhost:3000/test", nil // @@@
}

// SkillRoutineHandler handles the skill routine
func (req *RoutineRequest) SkillRoutineHandler(ctx *orva.SessionContext) {
	skillEndpoint, err := determineSkill(ctx.InitialInput.Message)
	if err != nil {
		return
	}

	skillResp, skillErr := orva.SkillProxy(skillEndpoint, ctx)

	// if an error occurs or if a skill response is not found
	// then we want the speech routine to kick in and handle the speech.
	if skillErr != nil {
		return
	}

	ctx.Append(skillResp)
}

// SpeechRoutineHandler handles the speech routine.
func (req *RoutineRequest) SpeechRoutineHandler(ctx *orva.SessionContext) {
	speechRequest := &grpcSpeech.SpeechRequest{
		Message: ctx.InitialInput.Message,
	}

	resp, err := req.SpeechClient.DetermineSpeech(context.Background(), speechRequest)
	if err != nil {
		resp := &orva.Response{
			Statement: "Having a hard time processing that",
		}

		ctx.Append(resp)
		return
	}

	mappedResp := mappings.SpeechToResponse(resp)
	ctx.Append(mappedResp)
	return
}
