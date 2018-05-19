package api

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	Err         string `json:"err,omitempty"`
}
