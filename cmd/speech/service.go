package main

import (
	"context"

	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc-speech"
)

// DetermineSpeech determines a speech statement given the request
func (server *ServiceRequest) DetermineSpeech(ctx *context.Context, *grpcSpeech.SpeechRequest) (*grpcSpeech.SpeechResponse, error) {
	return nil, nil
}

// CreateVariant creates a variant given an input speech statement
func (server *ServiceRequest) CreateVariant(ctx *context.Context, *grpcSpeech.SpeechRequest) (*grpcSpeech.SpeechRequest, error) {
	return nil, nil
}