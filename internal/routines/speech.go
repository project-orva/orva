package routines

import (
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"

	"github.com/GuyARoss/project-orva/pkg/orva"
)

// SpeechRoutineHandler handles the speech routine.
func (req *RoutineRequest) SpeechRoutineHandler(ctx *orva.SessionContext) {
	speechRequest := &grpcSpeech.SpeechRequest{
		Message: ctx.InitialInput.Message,
	}

	resp, err := req.SpeechClient.DetermineSpeech(nil, speechRequest)
	if err != nil {
		ctx.Append("Having a hard time processing that", "speech handler")
		return
	}

	ctx.Append(resp.Message, "speech handler")
	return
}
