package api

// ResetRequest - Reset the password
type ResetRequest struct {
	Username     string `json:"username"`
	RecoverToken string `json:"recover_token"` // contains email address
	Password     string `json:"password"`
}

// ResetResponse - Reset password response
type ResetResponse struct {
	Err string `json:"err,omitempty"`
}
