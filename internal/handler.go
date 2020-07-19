
package handler

import (
	"context"

	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	orva "github.com/GuyARoss/project-orva/pkg/orva"

	authService "github.com/GuyARoss/project-orva/pkg/services/auth"
	repoService "github.com/GuyARoss/project-orva/pkg/services/repository"
)
// RoutineRequest request used to interact with the handler
type RoutineRequest struct {
	SpeechClient grpcSpeech.GrpcSpeechClient
	SkillClient  grpcSkill.GrpcSkillClient

	AuthURI 	string 
	RepositoryURI string
}

// Invoke unifies orva routines
func (req *RoutineRequest) Invoke(ctx *orva.SessionContext) {
	user, err := req.validateRequest(ctx)
	if err != nil || user == nil {
		ctx.Append(&orva.Response{
			Statement: "Having trouble validating your request",
		})
		return
	}
	// @@ apply user to the routine handlers
	req.invokeRoutineHandlers(ctx)

	// @@ save routine ctx to long-term memory.
}

func (req *RoutineRequest) invokeRoutineHandlers(ctx *orva.SessionContext) {
	// @@ go routines.
	req.forwardContextToSkillService(ctx)
	req.forwardContextToSpeechService(ctx)

	if len(ctx.AppliedMessages) == 0 {
		ctx.Append(&orva.Response{
			Statement: "Having a hard time processing that",
		})
	}
}

func (req *RoutineRequest) forwardContextToSkillService(ctx *orva.SessionContext) {
	sq := &grpcSkill.ProcessRequest{
		Message: ctx.Request.Message,
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

// ForwardContextToSpeechService forward context to speech service
func (req *RoutineRequest) forwardContextToSpeechService(ctx *orva.SessionContext) {
	// @@ add the username to the request?
	sq := &grpcSpeech.SpeechRequest{
		Message: ctx.Request.Message,
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

func (req *RoutineRequest) validateRequest(ctx *orva.SessionContext) (*repoService.Profile, error) {
	dispatchPayload := &authService.DispatchAuthPayload{
		ID: ctx.Request.DeviceID,
		Key: ctx.Request.DeviceKey,
	}

	token, err := authService.FetchIdentityToken(req.AuthURI, dispatchPayload)
	if (err != nil || len(token.IdentityToken) == 0) {
		// cannot verify the devices origin, consider blocking or sending an alert out.
		return nil, err
	}
		
	// make call to profile repository with the identity token & prepare the user payload
	// if err then return invalid, else return valid.
	user, fetchErr := repoService.FetchUserProfile(
		req.RepositoryURI,
		ctx.Request.UserID,
		token.IdentityToken,
	)
	if fetchErr != nil {
		return nil, fetchErr
	}

	return user, nil
}
