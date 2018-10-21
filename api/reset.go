package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// ResetRequest - Reset the password
type ResetRequest struct {
	Username     string `json:"username"`
	RecoverToken string `json:"recover_token"` // contains email address
	Password     string `json:"password"`
}

// ResetResponse - Reset password response
type ResetResponse struct {
	Err string `json:"err,omitempty"`
}

// Reset - resets the password
func (a *AuthClient) Reset(req *ResetRequest) (*ResetResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/reset")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ResetResponse)
	return &response, nil
}
