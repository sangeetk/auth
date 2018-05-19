package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/model"
)

func (a Auth) Profile(ctx context.Context, req api.ProfileRequest) (api.ProfileResponse, error) {
	var response api.ProfileResponse
	var profile model.Profile
	var found = false

	if req.AccessToken != "" {
		identify := api.IdentifyRequest{AccessToken: req.AccessToken}
		u, err := a.Identify(ctx, identify)
		if err != nil {
			response.Err = err.Error()
			return response, nil
		}
		DB.Where("uid = ?", u.Uid).First(&profile)
		if profile.ID != 0 {
			found = true
		}
	}

	if found {
		return api.ProfileResponse{Profession: profile.Profession, Introduction: profile.Introduction}, nil
	}
	return api.ProfileResponse{Err: NotFound.Error()}, nil
}
