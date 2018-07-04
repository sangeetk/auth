package api

// ProfileRequest - Request the profile for the logged-in user
type ProfileRequest struct {
	AccessToken string `json:"access_token"`
}

// ProfileResponse - Returns the profile for the logged-in user
type ProfileResponse struct {
	Profile map[string]string `json:"profile,omitempty"`
	Err     string            `json:"err,omitempty"`
}
