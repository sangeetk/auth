package api

// ProfileRequest - Request the profile for the logged-in user
type ProfileRequest struct {
	AccessToken string `json:"access_token"`
}

// ProfileResponse - Returns the profile for the logged-in user
type ProfileResponse struct {
	Fields map[string]string `json:"fields"`
	Err    string            `json:"err,omitempty"`
}
