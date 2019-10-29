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

func EncodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoint.HashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}

func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return endpoint.HashRequest{Password: req.Password}, nil
}

func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.HashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return endpoint.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(endpoint.ValidateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return endpoint.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

func EncodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ValidateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}

func DecodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return endpoint.ValidateResponse{Valid: res.Valid}, nil
}

// NewGRPCServer gets a new pb.VaultServer.
func NewGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.VaultServer {
	return &grpcServer{
		hash: grpctransport.NewServer(
			endpoints.HashEndpoint,
			DecodeGRPCHashRequest,
			EncodeGRPCHashResponse,
		),
		validate: grpctransport.NewServer(
			endpoints.ValidateEndpoint,
			DecodeGRPCValidateRequest,
			EncodeGRPCValidateResponse,
		),
	}
}
