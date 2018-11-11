package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Domain    string   `json:"domain"`
	Roles     []string `json:"roles"`
	Err       string   `json:"err,omitempty"`
}

// Delete - deletes or deactivates the user
func Delete(req *DeleteRequest, dns string) (*DeleteResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/delete")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeDeleteResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(DeleteResponse)
	return &response, nil
}

// decodeDeleteResponse decodes the response from the service
func decodeDeleteResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response DeleteResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
