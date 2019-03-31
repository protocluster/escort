package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	envoy_auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
)

type AuthenticationService struct {
}

func (as *AuthenticationService) Check(ctx context.Context, req *envoy_auth.CheckRequest) (*envoy_auth.CheckResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
