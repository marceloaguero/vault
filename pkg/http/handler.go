package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/marceloaguero/vault/pkg/endpoint"
)

func decodeHashRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.HashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
