package endpoint

import (
	"context"
	"errors"

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
func MakeHashEndpoint(srv service.Vault) endpoint.Endpoint {
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
func MakeValidateEndpoint(srv service.Vault) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return ValidateResponse{false, err.Error()}, nil
		}

		return ValidateResponse{v, ""}, nil
	}
}

// Endpoints represents all endpoints for the vault Service.
type Endpoints struct {
	HashEndpoint     endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

// Implement service.Vault interface

// Hash uses the HashEndpoint to hash a password.
func (e Endpoints) Hash(ctx context.Context, password string) (string, error) {
	req := HashRequest{Password: password}
	resp, err := e.HashEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	hashResp := resp.(HashResponse)
	if hashResp.Err != "" {
		return "", errors.New(hashResp.Err)
	}
	return hashResp.Hash, nil
}

// Validate uses the ValidateEndpoint to validate a password and hash pair.
func (e Endpoints) Validate(ctx context.Context, password, hash string) (bool, error) {
	req := ValidateRequest{Password: password, Hash: hash}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(ValidateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
