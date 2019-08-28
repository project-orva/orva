package routines

import grpcProfile "github.com/GuyARoss/project-orva/pkg/grpc/profile"

type RoutineRequest struct {
	ProfileClient grpcProfile.GrpcProfileClient
}
