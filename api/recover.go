package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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
func Recover(req *RecoverRequest, dns string) (*RecoverResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/recover")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeRecoverResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(RecoverResponse)
	return &response, nil
}

// decodeRecoverResponse decodes the response from the service
func decodeRecoverResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response RecoverResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
