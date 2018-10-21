package api

import (
	"context"
	"log"
	"net/url"

	ht "github.com/urantiatech/kit/transport/http"
)

// DeleteRequest - Delete the existing user
type DeleteRequest struct {
	AccessToken string `json:"access_token"`
	Password    string `json:"password"`
}

// DeleteResponse - Returns error if delete fails
type DeleteResponse struct {
	Err string `json:"err,omitempty"`
}

// Delete - deletes or deactivates the user
func (a *AuthClient) Delete(req *DeleteRequest) (*DeleteResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/delete")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(DeleteResponse)
	return &response, nil
}
