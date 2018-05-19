package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
)

func (Auth) Refresh(_ context.Context, req api.RefreshRequest) (api.RefreshResponse, error) {
	var response api.RefreshResponse
	return response, nil
}
