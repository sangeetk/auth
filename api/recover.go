package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// RecoverRequest - Recover the account
type RecoverRequest struct {
	Username string `json:"username"`
}

// RecoverResponse - Reset the new password
type RecoverResponse struct {
	RecoverToken string `json:"recover_token,omitempty"`
	Err          string `json:"err,omitempty"`
}

// Recover - sends a recovery mail to reset the password
func (a *AuthClient) Recover(req *RecoverRequest) (*RecoverResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/recover")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(RecoverResponse)
	return &response, nil
}
