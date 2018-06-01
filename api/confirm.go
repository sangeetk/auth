package api

// ConfirmRequest - Confirms the new registration
type ConfirmRequest struct {
	Username     string `json:"username"`
	ConfirmToken string `json:"confirm_token"`
}

// ConfirmResponse - Returns error if registration confirmation fails
type ConfirmResponse struct {
	Err string `json:"err,omitempty"`
}
