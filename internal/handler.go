
package handler

import (
	"context"
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
	"encoding/json"

	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	orva "github.com/GuyARoss/project-orva/pkg/orva"
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
	if req.SkillClient == nil {
		return
	}

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
	if req.SpeechClient == nil {
		return
	}

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

type DispatchAuthPayload struct {
	ID string `json:"resource_id"`
	Key string `json:"resource_key"`	
}

type DispatchAuthResponse struct {
	IdentityToken string `json:"identity_token"`
	IAT uint64 `json:"iat"`
}

func (req *RoutineRequest) fetchIdentityToken(ctx *orva.SessionContext) (*DispatchAuthResponse, error) {
	dispatchPayload := &DispatchAuthPayload{
		ID: ctx.Request.DeviceID,
		Key: ctx.Request.DeviceKey,
	}

	payloadBytes, marshalErr := json.Marshal(dispatchPayload)
	if marshalErr != nil {
		return nil, marshalErr
	}

	authURI := fmt.Sprintf("%s/dispatch", req.AuthURI)
	authRequest, err := http.NewRequest("POST", authURI, bytes.NewBuffer(payloadBytes))
	authRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(authRequest)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ioout, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return nil, ioErr
	}

	dispatchResp := &DispatchAuthResponse{}
	marErr := json.Unmarshal(ioout, dispatchResp)
	if marErr != nil {
		return nil, marErr
	}

	return dispatchResp, nil
}

type User struct {
	FirstName string `json:"firstName"`
	lastName string `json:"lastName"`
	AccessLevel uint32 `json:"accessLevel"`
	ID string `json:"id"`
}

func (req *RoutineRequest) fetchUserProfile(input *orva.SessionRequest, identityToken string) (*User, error) {
	uri := fmt.Sprintf("%s/profile?id=%s", req.RepositoryURI, input.UserID)
	httpReq, err := http.NewRequest("GET", uri, nil)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-identity", identityToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ioout, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return nil, ioErr
	}

	dispatchResp := &User{}
	marErr := json.Unmarshal(ioout, dispatchResp)
	if marErr != nil {
		return nil, marErr
	}

	return dispatchResp, nil
}

func (req *RoutineRequest) validateRequest(ctx *orva.SessionContext) (*User, error) {
	token, err := req.fetchIdentityToken(ctx)
	if err != nil {
		return nil, err
	}

	// make call to profile repository with the identity token & prepare the user payload
	// if err then return invalid, else return valid.
	user, fetchErr := req.fetchUserProfile(&ctx.Request, token.IdentityToken)
	if fetchErr != nil {
		return nil, fetchErr
	}

	return user, nil
}
