package api

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

type LogoutResponse struct {
	Err string `json:"err,omitempty"`
}
