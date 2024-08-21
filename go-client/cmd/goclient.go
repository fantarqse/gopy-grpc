package goclientcmd

import (
	"go-client/internal/grpc"
	"go-client/internal/http"
)

func Run() error {
	// INFO: probably, server or client should run in a separated goroutine
	grpcClient := grpc.NewClient()
	if err := grpcClient.Serve(); err != nil {
		return err
	}

	httpServer := http.NewServer()
	if err := httpServer.Serve("5005"); err != nil {
		return err
	}

	return nil
}
