package main

import (
	pb "binnacle-grpc-server/proto/trip"
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = ":8080"
)

type healthChecker struct{}

func (h healthChecker) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h healthChecker) Watch(in *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

func main() {
	port := flag.String("port", "", "port to listen on")
	flag.Parse()

	if *port == "" {
		log.Println("No port specified, using default port 8080")
		*port = defaultPort
	}

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthChecker{})
	reflection.Register(s)

	log.Println("Starting gRPC server on port 8080")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
