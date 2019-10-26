package main

import (
	"context"

	coreHandler "github.com/GuyARoss/project-orva/internal/core_handler"
	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
	"github.com/GuyARoss/project-orva/pkg/orva"
	"github.com/GuyARoss/project-orva/pkg/utilities"
)

type ServiceRequest struct {
	*coreHandler.RoutineRequest
}

// ProcessStatement processes the inputted statement, outputs a response.
func (server *ServiceRequest) ProcessStatement(ctx context.Context, req *grpcCore.Request) (*grpcCore.Response, error) {
	t := &utilities.TimeStamp{}
	t.Start()

	input := &orva.Input{Message: req.Message, DeviceID: req.DeviceID, UserID: req.UserID}
	octx := orva.CreateContext(input)

	server.RoutineRequest.CoreHandler(octx)
	responses := octx.Responses()
	derefResponses := *responses

	// responses -> grpcCore::Response
	contextResponse := make([]*grpcCore.ContextResponse, len(derefResponses))
	for i := 0; i < len(derefResponses); i++ {
		contextResponse[i] = &grpcCore.ContextResponse{
			Message:      derefResponses[i].Statement,
			AssignedFrom: derefResponses[i].AssignedFrom,
			GraphicURL:   derefResponses[i].GraphicURL,
			GraphicType:  derefResponses[i].GraphicType,
			Error:        derefResponses[i].Error,
		}
	}

	t.Stop()

	return &grpcCore.Response{
		Duration:        int32(t.Distance()),
		ContextResponse: contextResponse,
	}, nil
}
