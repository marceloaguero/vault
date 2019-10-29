package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/marceloaguero/vault/pb"
	"github.com/marceloaguero/vault/pkg/endpoint"
)

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}

func (s *grpcServer) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ValidateResponse), nil
}

func encodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoint.HashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}

func decodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return endpoint.HashRequest{Password: req.Password}, nil
}

func encodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.HashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

func decodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return endpoint.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

func encodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoint.ValidateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

func decodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return endpoint.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

func encodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ValidateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}

func decodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return endpoint.ValidateResponse{Valid: res.Valid}, nil
}

// NewGRPCServer gets a new pb.VaultServer.
func NewGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.VaultServer {
	return &grpcServer{
		hash: grpctransport.NewServer(
			endpoints.HashEndpoint,
			decodeGRPCHashRequest,
			encodeGRPCHashResponse,
		),
		validate: grpctransport.NewServer(
			endpoints.ValidateEndpoint,
			decodeGRPCValidateRequest,
			encodeGRPCValidateResponse,
		),
	}
}
