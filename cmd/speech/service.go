package main

import (
	"context"

	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
)

type ServiceRequest struct {
}


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
		GraphicType: 0,
		Error: "",
	}, nil
}

// CreateVariant creates a variant given an input speech statement
func (server *ServiceRequest) CreateVariant(
	ctx context.Context,
	req *grpcSpeech.SpeechRequest,
) (
	*grpcSpeech.SpeechResponse,
	error,
) {
	return &grpcSpeech.SpeechResponse{
		Duration: 0,
		Message: req.Message,
		GraphicURL: "",
		GraphicType: 0,
		Error: "",	
	}, nil
}