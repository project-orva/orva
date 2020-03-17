package main

import (
	"fmt"
	"testing"

	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	"github.com/GuyARoss/project-orva/pkg/orva"
)

func TestSpeechRoutineHandler(t *testing.T) {
	req := &RoutineRequest{
		SpeechClient: grpcSpeech.CreateClientConn("nuffn"),
		PgCreds:      nil,
	}

	ctx := &orva.SessionContext{
		InitialInput: &orva.Input{
			Message:  "Does this work?",
			DeviceID: "testdevice",
			UserID:   "testuser",
		},
	}

	req.SpeechRoutineHandler(ctx)

	f := *ctx.Responses()
	if len(f) != 1 {
		t.Error("dis thing should resolve")
	}
	fmt.Println(*ctx.Responses())
}
