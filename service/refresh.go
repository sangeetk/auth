package service

import (
	"context"

	"git.urantiatech.com/auth/auth/api"
)

// Refresh - Generates a new token and extends existing session
func (Auth) Refresh(_ context.Context, req api.RefreshRequest) (api.RefreshResponse, error) {
	var response api.RefreshResponse
	return response, nil
}
