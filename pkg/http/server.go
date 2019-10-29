package http

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/marceloaguero/vault/pkg/endpoint"
)

// NewHTTPServer makes a new Vault HTTP service
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hash", httptransport.NewServer(
		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))
	m.Handle("/validate", httptransport.NewServer(
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))
	return m
}
