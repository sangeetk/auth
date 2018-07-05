package api

// AuthorizationRequest - Authorize against the role
type AuthorizationRequest struct {
	AccessToken string `json:"access_token"`
	Domain      string `json:"domain"`
	Role        string `json:"role"`
}

// AuthorizationResponse - Returns Token or  error
type AuthorizationResponse struct {
	Authorize bool   `json:"authorize"`
	Err       string `json:"err,omitempty"`
}
