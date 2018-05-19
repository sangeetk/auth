package api

type ConfirmRequest struct {
	ConfirmToken string `json:"confirm_token"`
}

type ConfirmResponse struct {
	Err string `json:"err,omitempty"`
}
