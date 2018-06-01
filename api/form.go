package api

// Request - Generic request
type Request struct {
	URI    string `json:"uri"`
	Domain string `json:"domain"`
	Data   string `json:"data"`
}

// Response - Generic response
type Response struct {
	HTML        string `json:"html"`
	AccessToken string `json:"access_token,omitempty"`
	Err         string `json:"err,omitempty"`
}
