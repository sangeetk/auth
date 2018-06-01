package api

// LogoutRequest - Closes the session
type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

// LogoutResponse - Returns error for invalid sessions
type LogoutResponse struct {
	Err string `json:"err,omitempty"`
}
