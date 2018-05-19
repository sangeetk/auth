package api

type Request struct {
	Uri    string `json:"uri"`
	Domain string `json:"domain"`
	Data   string `json:"data"`
}

type Response struct {
	Html        string `json:"html"`
	AccessToken string `json:"access_token,omitempty"`
	Err         string `json:"err,omitempty"`
}
