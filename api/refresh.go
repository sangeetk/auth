package api

type RefreshRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	NewAccessToken string `json:"new_access_token"`
	Err            string `json:"err,omitempty"`
}
