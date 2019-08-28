package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc-skill"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func main() {
	tcpPort := flag.String("p", "3005", "specified port to start the gRPC server on")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *tcpPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		PermitWithoutStream: true,
		MinTime:             1 * time.Minute,
	}))

	grpcSkill.RegisterGrpcSkillServer(grpcServer, &ServiceRequest{})
	reflection.Register(grpcServer)

	fmt.Println(fmt.Sprintf("gRPC service started on port %s", *tcpPort))
	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}

	grpcServer.Serve(listener)
}
