package api

type ProfileRequest struct {
	AccessToken string `json:"access_token"`
}

type ProfileResponse struct {
	Profession   string `json:"profession"`
	Introduction string `json:"introduction"`
	Err          string `json:"err,omitempty"`
}
