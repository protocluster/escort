package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	envoy "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

type ClusterDiscoveryService struct {
}

func (cds *ClusterDiscoveryService) StreamClusters(call envoy.ClusterDiscoveryService_StreamClustersServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (cds *ClusterDiscoveryService) DeltaClusters(call envoy.ClusterDiscoveryService_DeltaClustersServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (cds *ClusterDiscoveryService) FetchClusters(ctx context.Context, req *envoy.DiscoveryRequest) (*envoy.DiscoveryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
