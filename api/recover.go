package api

const (
	RecoveryToken = "recover"
	ResetPassword = "reset"
)

type RecoverRequest struct {
	// Should be set to either ResetPassword or RecoveryToken
	Cmd string `json:"cmd"`

	// For RecoveryToken request
	Email string `json:"email"`

	// For ResetPassword request
	RecoverToken string `json:"recover_token"` // contains email address
	Password     string `json:"password"`
}

type RecoverResponse struct {
	RecoverToken string `json:"recover_token,omitempty"`
	Err          string `json:"err,omitempty"`
}
