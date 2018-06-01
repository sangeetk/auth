package api

const (
	// RecoveryToken -
	RecoveryToken = "recover"
	// ResetPassword -
	ResetPassword = "reset"
)

// RecoverRequest - Reset the password
type RecoverRequest struct {
	// Should be set to either ResetPassword or RecoveryToken
	Cmd string `json:"cmd"`

	// For RecoveryToken request
	Username string `json:"username"`

	// For ResetPassword request
	RecoverToken string `json:"recover_token"` // contains email address
	Password     string `json:"password"`
}

// RecoverResponse - Reset the new password
type RecoverResponse struct {
	RecoverToken string `json:"recover_token,omitempty"`
	Err          string `json:"err,omitempty"`
}
