package main

import (
	"context"

	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
)

// DetermineSpeech determines a speech statement given the request
func (server *ServiceRequest) DetermineSpeech(
	ctx context.Context,
	req *grpcSpeech.SpeechRequest,
) (
	*grpcSpeech.SpeechResponse,
	error,
) {
	return &grpcSpeech.SpeechResponse{
		Duration: 0,
		Message: "",
		GraphicURL: "",
		GraphicType: "",
		Error: "",
	}, nil
}

// CreateVariant creates a variant given an input speech statement
func (server *ServiceRequest) CreateVariant(
	ctx context.Context,
	req *grpcSpeech.SpeechRequest
) (
	*grpcSpeech.SpeechRequest,
	error,
) {
	return &grpcSpeech.SpeechResponse{
		Duration: 0,
		Message: *req.Message,
		GraphicURL: "",
		GraphicType: "",
		Error: "",	
	}, nil
}