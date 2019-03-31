package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	envoy "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

var (
	port = flag.Int("port", 5333, "gRPC listen port")
)

// https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/http/jwt_authn/v2alpha/config.proto
// https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/lds.proto
// https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/ext_authz_filter

func main() {
	address := fmt.Sprintf(":%d", *port)
	listener, err := net.Listen("tcp", address)

	server := grpc.NewServer()
	cds := &ClusterDiscoveryService{}
	envoy.RegisterClusterDiscoveryServiceServer(server, cds)

	rds := &RouteDiscoveryService{}
	envoy.RegisterRouteDiscoveryServiceServer(server, rds)

	as := &AuthenticationService{}
	envoy_auth.RegisterAuthorizationServer(server, as)

	log.Printf("xds listenting on %s", address)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
