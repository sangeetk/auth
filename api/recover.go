package api

// RecoverRequest - Recover the account
type RecoverRequest struct {
	Username string `json:"username"`
}

// RecoverResponse - Reset the new password
type RecoverResponse struct {
	RecoverToken string `json:"recover_token,omitempty"`
	Err          string `json:"err,omitempty"`
}
