package grpcclient

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/marceloaguero/vault/pb"
	"github.com/marceloaguero/vault/pkg/endpoint"
	grpcservice "github.com/marceloaguero/vault/pkg/grpc"
	"github.com/marceloaguero/vault/pkg/service"
	"google.golang.org/grpc"
)

// New makes a new vault.Service client.
func New(conn *grpc.ClientConn) service.VaultService {
	var hashEndpoint = grpctransport.NewClient(
		conn, "Vault", "Hash",
		grpcservice.EncodeGRPCHashRequest,
		grpcservice.DecodeGRPCHashResponse,
		pb.HashResponse{},
	).Endpoint()
	var validateEndpoint = grpctransport.NewClient(
		conn, "Vault", "Validate",
		grpcservice.EncodeGRPCValidateRequest,
		grpcservice.DecodeGRPCValidateResponse,
		pb.ValidateResponse{},
	).Endpoint()
	return endpoint.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
}
