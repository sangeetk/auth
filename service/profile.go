package service

import (
	"context"
	"encoding/json"
	"net/http"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/go-kit/kit/endpoint"
)

// Profile - Returns user profile
func (a Auth) Profile(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	var response = &api.ProfileResponse{}

	t, err := token.ParseToken(req.AccessToken)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Check against Blacklisted tokens
	if _, found := token.BlacklistAccessTokens.Get(req.AccessToken); found {
		response.Err = token.ErrorInvalidToken.Error()
		return response, nil
	}

	// Get user details
	u, err := user.Read(t.Username)
	if err != nil || !u.Confirmed || !u.DeletedAt.IsZero() {
		response.Err = ErrorInvalidLogin.Error()
		return response, nil
	}

	// Add fields
	response.Username = u.Username
	response.Name = u.Name
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Birthday = u.Birthday
	response.InitialDomain = u.InitialDomain
	response.Roles = u.Roles
	response.Address = u.Address
	response.Profile = u.Profile

	return response, nil
}

// MakeProfileEndpoint -
func MakeProfileEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.ProfileRequest)
		return svc.Profile(ctx, &req)
	}
}

// DecodeProfileRequest -
func DecodeProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.ProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
