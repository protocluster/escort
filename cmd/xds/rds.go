package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	envoy "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

type RouteDiscoveryService struct {
}

func (rds *RouteDiscoveryService) StreamRoutes(call envoy.RouteDiscoveryService_StreamRoutesServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (rds *RouteDiscoveryService) DeltaRoutes(envoy.RouteDiscoveryService_DeltaRoutesServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (rds *RouteDiscoveryService) FetchRoutes(ctx context.Context, req *envoy.DiscoveryRequest) (*envoy.DiscoveryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
