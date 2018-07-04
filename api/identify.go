package api

// IdentifyRequest - Indentifies the user using JWT
type IdentifyRequest struct {
	AccessToken string `json:"access_token"`
}

// IdentifyResponse - Sends the user details as reponse
type IdentifyResponse struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Err       string `json:"err,omitempty"`
}
