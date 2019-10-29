package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/marceloaguero/vault/pb"
	"github.com/marceloaguero/vault/pkg/endpoint"
	grpcservice "github.com/marceloaguero/vault/pkg/grpc"
	httpservice "github.com/marceloaguero/vault/pkg/http"
	"github.com/marceloaguero/vault/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":50051", "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := service.NewVaultService()
	errChan := make(chan error)

	// Trap termination signals (such as Ctrl+C) and send an error to errChan
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// Create endpoints
	hashEndpoint := endpoint.MakeHashEndpoint(srv)
	validateEndpoint := endpoint.MakeValidateEndpoint(srv)
	endpoints := endpoint.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}

	// HTTP transport
	go func() {
		log.Println("http:", *httpAddr)
		handler := httpservice.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// gRPC transport
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := grpcservice.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	// Prevent main func to exit.
	// Block it waiting until something tells the program to terminate.
	// Just listen errChan. It will block while nothing is sent to it.
	log.Fatalln(<-errChan)
}
