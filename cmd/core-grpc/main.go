package main

import (
	"os"
	"flag"
	"fmt"
	"net"
	"time"

	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	grpcSkill "github.com/GuyARoss/project-orva/pkg/grpc/skill"
	handler "github.com/GuyARoss/project-orva/internal"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func main() {
	tcpPort := flag.String("p", "3005", "specified port to start the gRPC server on")
	envErr := godotenv.Load("../../.env")
	if envErr != nil {
		panic(envErr)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *tcpPort))
	if err != nil {
		panic(err)
	}

	rr := &handler.RoutineRequest{
		SkillClient: grpcSkill.CreateClientConn(os.Getenv("SKILL_URI")),
		SpeechClient: grpcSpeech.CreateClientConn(os.Getenv("SPEECH_URI")),
		AuthURI: os.Getenv("AUTH_URI"),
		RepositoryURI: os.Getenv("REPO_URI"),
	}

	grpcServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		PermitWithoutStream: true,
		MinTime:             1 * time.Minute,
	}))

	// grpcProfile.RegisterGrpcProfileServer(grpcServer, &ServiceRequest{})
	grpcCore.RegisterGrpcCoreServer(grpcServer, &ServiceRequest{
		RoutineRequest: rr,		
	})
	reflection.Register(grpcServer)

	fmt.Println(fmt.Sprintf("gRPC service started on port %s", *tcpPort))
	if serveErr := grpcServer.Serve(listener); err != nil {
		panic(serveErr)
	}

	grpcServer.Serve(listener)
}
