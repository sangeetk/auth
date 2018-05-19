package api

type IdentifyRequest struct {
	AccessToken string `json:"access_token"`
}

type IdentifyResponse struct {
	Uid   uint64      `json:"uid"`
	Fname string      `json:"fname"`
	Lname string      `json:"lname"`
	Email string      `json:"email"`
	Roles interface{} `json:"roles"`
	Err   string      `json:"err,omitempty"`
}
