package main

import (
	"context"

	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
	"github.com/GuyARoss/project-orva/pkg/orva"
	"github.com/GuyARoss/project-orva/pkg/utilities"
)

type ServiceRequest struct {
	RoutineRequest *RoutineRequest
}

// ProcessStatement processes the inputted statement, outputs a response.
func (server *ServiceRequest) ProcessStatement(ctx context.Context, req *grpcCore.Request) (*grpcCore.Response, error) {
	t := &utilities.TimeStamp{}
	t.Start()

	input := &orva.Input{Message: req.Message, DeviceID: req.DeviceID, UserID: req.UserID}
	octx := orva.CreateContext(input)

	server.RoutineRequest.CoreHandler(octx)

	responses := octx.Responses()
	rlen := len(responses)

	// responses -> grpcCore::Response
	contextResponses := make([]*grpcCore.Response_ContextResponse, rlen)
	for i := 0; i < rlen; i++ {
		contextResponses[i] = &grpcCore.Response_ContextResponse{
			Message:  responses[i].Statement,
			Gph:      responses[i].GraphicURL,
			GphType:  responses[i].GraphicType,
			Duration: responses[i].Duration,
			// @@ remove the error from the protobuf
		}
	}

	t.Stop()

	return &grpcCore.Response{
		Duration:         float32(t.Distance()),
		ContextResponses: contextResponses,
	}, nil
}
