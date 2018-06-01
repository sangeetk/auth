package api

// AuthorizationRequest - Authorize agains the role
type AuthorizationRequest struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
}

// AuthorizationResponse - Returns Unauthrize on error
type AuthorizationResponse struct {
	AuthToken string `json:"auth_token"`
	Err       string `json:"err,omitempty"`
}
