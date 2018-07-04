package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/user"
)

// Profile - Returns user profile
func (a Auth) Profile(ctx context.Context, req api.ProfileRequest) (api.ProfileResponse, error) {
	var response api.ProfileResponse
	var u user.User
	var found = false

	if req.AccessToken != "" {
		identify := api.IdentifyRequest{AccessToken: req.AccessToken}
		user, err := a.Identify(ctx, identify)
		if err != nil {
			response.Err = err.Error()
			return response, nil
		}
		found = true
		u.Username = user.Username
	}

	if found {
		u, err := user.Read(u.Username)
		if err != nil {
			response.Err = ErrorNotFound.Error()
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
	return api.ProfileResponse{Err: ErrorNotFound.Error()}, nil
}
