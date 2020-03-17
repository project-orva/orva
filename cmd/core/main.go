package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
	grpcSpeech "github.com/GuyARoss/project-orva/pkg/grpc/speech"
	"github.com/GuyARoss/project-orva/pkg/pgdb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func main() {
	tcpPort := flag.String("p", "3005", "specified port to start the gRPC server on")
	pgAddress := flag.String("pg", "http://localhost:32768", "postgresql database address")
	speechAdress := flag.String("s", "http://localhost:5521", "speech service address")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *tcpPort))
	if err != nil {
		panic(err)
	}

	rr := &RoutineRequest{
		SpeechClient: grpcSpeech.CreateClientConn(*speechAdress),
		PgCreds: &pgdb.DbCreds{ // @@@ nut done, plz fix
			Host:     *pgAddress,
			User:     "docker",
			Password: "docker",
			Dbname:   "docker",
		},
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
