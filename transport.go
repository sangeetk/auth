package main

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/service"
	"github.com/urantiatech/kit/endpoint"
)

func makeRegisterEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RegisterRequest)
		return svc.Register(ctx, req)
	}
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeUpdateEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.UpdateRequest)
		return svc.Update(ctx, req)
	}
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeConfirmEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ConfirmRequest)
		return svc.Confirm(ctx, req)
	}
}

func decodeConfirmRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ConfirmRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeLoginEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LoginRequest)
		return svc.Login(ctx, req)
	}
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeLogoutEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LogoutRequest)
		return svc.Logout(ctx, req)
	}
}

func decodeLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeIdentifyEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.IdentifyRequest)
		return svc.Identify(ctx, req)
	}
}

func decodeIdentifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.IdentifyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeProfileEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ProfileRequest)
		return svc.Profile(ctx, req)
	}
}

func decodeProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeRefreshEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RefreshRequest)
		return svc.Refresh(ctx, req)
	}
}

func decodeRefreshRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeRecoverEndpoint(svc service.Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RecoverRequest)
		return svc.Recover(ctx, req)
	}
}

func decodeRecoverRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RecoverRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
