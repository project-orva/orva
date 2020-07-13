package handler

import (
	"context"

	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	"github.com/GuyARoss/project-orva/pkg/orva"
)

// RoutineRequest request used to interact with the handler
type RoutineRequest struct {
	SpeechClient grpcSpeech.GrpcSpeechClient
	SkillClient  grpcSkill.GrpcSkillClient

	AuthURI 	string 
}

// CoreHandler unifies orva routines
func (req *RoutineRequest) Invoke(ctx *orva.SessionContext) {
	err := req.validateRequest()
	if err != nil {
		ctx.Append(&orva.Response{
			Statement: "Having trouble validating your request",
		})
		return
	}

	req.invokeRoutineHandlers(ctx)

	// @@ save routine ctx to long-term memory.
}

func (req *RoutineRequest) invokeRoutineHandlers(ctx *orva.SessionContext) {
	// @@ go routines.
	req.fowardContextToSkillService(ctx)
	req.fowardContextToSpeechService(ctx)

	if len(ctx.AppliedMessages) == 0 {
		ctx.Append(&orva.Response{
			Statement: "Having a hard time processing that",
		})
	}
}

func (req *RoutineRequest) fowardContextToSkillService(ctx *orva.SessionContext) {
	if req.SkillClient == nil {
		return
	}

	sq := &grpcSkill.ProcessRequest{
		Message: ctx.InitialInput.Message,
		TransactionID: "test123",
	}
	
	resp, err := req.SkillClient.ProcessSkillRequest(context.Background(), sq)
	if err != nil || len(resp.ForwardAddress) < 0 {
		return
	}

	skillResp, skillErr := orva.SkillProxy(resp.ForwardAddress, ctx)

	if skillErr == nil {
		ctx.Append(skillResp)
		return
	}
}

// FowardContextToSpeechService fowards context to speech service
func (req *RoutineRequest) fowardContextToSpeechService(ctx *orva.SessionContext) {
	if req.SpeechClient == nil {
		return
	}

	// @@ add the username to the request?
	sq := &grpcSpeech.SpeechRequest{
		Message: ctx.InitialInput.Message,
	}

	resp, err := req.SpeechClient.HandleSpeechRequest(context.Background(), sq)
	if err != nil {
		return
	}

	ctx.Append(&orva.Response{
		Statement:   resp.Message,
		GraphicURL:  resp.GraphicURL,
		GraphicType: resp.GraphicType,
	})
}

// AccountRoutineHandler handles the profiles routine
func (req *RoutineRequest) accountRoutineHandler(ctx *orva.SessionContext) {

}

func (req *RoutineRequest) validateRequest() error {
	// @@ send the device id and key to auth service
	// get back the token
	// make call to profile respository with the identity token & prepare the user payload
	// if err then return invalid, else return valid.
	return nil
}
