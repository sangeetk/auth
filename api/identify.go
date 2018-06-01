package api

// IdentifyRequest - Indentifies the user using JWT
type IdentifyRequest struct {
	AccessToken string `json:"access_token"`
}

// IdentifyResponse - Sends the user details as reponse
type IdentifyResponse struct {
	Username  string      `json:"username"`
	FirstName string      `json:"fname"`
	LastName  string      `json:"lname"`
	Email     string      `json:"email"`
	Roles     interface{} `json:"roles"`
	Err       string      `json:"err,omitempty"`
}
