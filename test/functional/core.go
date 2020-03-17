package main

import (
	"context"
	"fmt"

	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
)

func main() {
	r := grpcCore.CreateClientConn("localhost:3005")

	req := &grpcCore.Request{
		Message:  "this working?",
		UserID:   "test_account",
		DeviceID: "test_device",
	}

	response, err := r.ProcessStatement(context.Background(), req)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
