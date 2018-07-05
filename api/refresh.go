package api

// RefreshRequest - Request to refresh the JWT token
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshResponse - Returns new JWT token on success
type RefreshResponse struct {
	NewAccessToken  string `json:"new_access_token,omitempty"`
	NewRefreshToken string `json:"new_refresh_token,omitempty"`
	Err             string `json:"err,omitempty"`
}
