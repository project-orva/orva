package main

import "github.com/GuyARoss/project-orva/pkg/settings"

// ServiceRequest cmd service request required by grpc
type ServiceRequest struct {
	PostgresSettings *settings.AWSRDS
}
