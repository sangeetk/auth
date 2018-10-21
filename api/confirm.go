package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// ConfirmRequest - Confirms the new registration
type ConfirmRequest struct {
	Username     string `json:"username"`
	ConfirmToken string `json:"confirm_token"`
}

// ConfirmResponse - Returns error if registration confirmation fails
type ConfirmResponse struct {
	Err string `json:"err,omitempty"`
}

// Confirm - confirms the new user registration
func (a *AuthClient) Confirm(req *ConfirmRequest) (*ConfirmResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/confirm")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ConfirmResponse)
	return &response, nil
}
