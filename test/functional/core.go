package main

import (
	"context"
	"fmt"

	grpcCore "github.com/GuyARoss/project-orva/pkg/grpc/core"
)

func main() {
	r := grpcCore.CreateClientConn("localhost:3003")

	req := &grpcCore.Request{
		Message:  "this working?",
		UserID:   "test_account",
		DeviceID: "test_device",
		DeviceKey: "47344117-8fd8-4e80-bcf5-4e0c10c8d14d",
	}

	response, err := r.ProcessStatement(context.Background(), req)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
