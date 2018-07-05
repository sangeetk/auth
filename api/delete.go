package api

// DeleteRequest - Delete the existing user
type DeleteRequest struct {
	AccessToken string `json:"access_token"`
	Password    string `json:"password"`
}

// DeleteResponse - Returns error if delete fails
type DeleteResponse struct {
	Err string `json:"err,omitempty"`
}
