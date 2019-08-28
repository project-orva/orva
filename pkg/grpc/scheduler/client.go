package grpcScheduler

import (
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// CreateClientConn creates a new core client connection
func CreateClientConn(hostAddress string) GrpcSchedulerClient {
	k := keepalive.ClientParameters{
		Time:                time.Minute * 2,
		PermitWithoutStream: true,
	}

	cc, err := grpc.Dial(hostAddress, grpc.WithInsecure(), grpc.WithKeepaliveParams(k))
	if err != nil {
		panic(err)
	}

	return NewGrpcSchedulerClient(cc)
}
