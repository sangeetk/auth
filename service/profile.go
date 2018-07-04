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
		response.Profile = make(map[string]string)
		// Add fields
		response.Profile["profession"] = u.Profile.Profession
		response.Profile["introduction"] = u.Profile.Introduction

		return response, nil
	}
	return api.ProfileResponse{Err: ErrorNotFound.Error()}, nil
}
