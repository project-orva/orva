package main

import (
	"context"

	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	"github.com/GuyARoss/project-orva/pkg/orva"
	pgDevice "github.com/GuyARoss/project-orva/pkg/pg-schemas/device"
	pgUser "github.com/GuyARoss/project-orva/pkg/pg-schemas/user"
	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

// RoutineRequest request used to interact with the handler
type RoutineRequest struct {
	SpeechClient grpcSpeech.GrpcSpeechClient
	PgCreds      *pgdb.DbCreds
}

// CoreHandler unifies orva routines
func (req *RoutineRequest) CoreHandler(ctx *orva.SessionContext) {
	req.invokeRoutineHandlers(ctx)

	// @@ save routine ctx to long-term memory.
}

func (req *RoutineRequest) invokeRoutineHandlers(ctx *orva.SessionContext) {
	// @@ make these go routines.
	req.accountRoutineHandler(ctx)
	req.fowardContextToSpeechService(ctx)
}

// FowardContextToSpeechService fowards context to speech service
func (req *RoutineRequest) fowardContextToSpeechService(ctx *orva.SessionContext) {
	// @@ add the username to the request?
	sq := &grpcSpeech.SpeechRequest{
		Message: ctx.InitialInput.Message,
	}

	resp, err := req.SpeechClient.HandleSpeechRequest(context.Background(), sq)
	if err != nil {
		resp := &orva.Response{
			Statement: "Having a hard time processing that",
		}

		ctx.Append(resp)
		return
	}

	// foward to skill proxy
	if len(resp.FowardAddress) > 0 {
		skillResp, skillErr := orva.SkillProxy(resp.FowardAddress, ctx)

		if skillErr != nil {
			// @@ response that we are having trouble processing dat?
			return
		}
		ctx.Append(skillResp)
	}

	ctx.Append(&orva.Response{
		Statement:   resp.Message,
		Duration:    resp.Duration,
		GraphicURL:  resp.GraphicURL,
		GraphicType: resp.GraphicType,
	})
}

// AccountRoutineHandler handles the profiles routine
func (req *RoutineRequest) accountRoutineHandler(ctx *orva.SessionContext) {
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
