package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/marceloaguero/vault/pkg/service"
)

// HashRequest specifies the request parameters for Hash method
type HashRequest struct {
	Password string `json:"password"`
}

// HashResponse specifies the response parameters for Hash method
type HashResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

// ValidateRequest specifies the request parameters for validate method
type ValidateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

// ValidateResponse specifies the response parameters for Validate method
type ValidateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

// MakeHashEndpoint returns an endpoint that invokes Hash on the service
func MakeHashEndpoint(srv service.VaultService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HashRequest)
		v, err := srv.Hash(ctx, req.Password)
		if err != nil {
			return HashResponse{v, err.Error()}, nil
		}

		return HashResponse{v, ""}, nil
	}
}

// MakeValidateEndpoint returns an endpoint that invokes Validate on the service
func MakeValidateEndpoint(srv service.VaultService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return ValidateResponse{false, err.Error()}, nil
		}

		return ValidateResponse{v, ""}, nil
	}
}
