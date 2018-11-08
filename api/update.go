package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	ht "github.com/urantiatech/kit/transport/http"
)

// UpdateRequest - Update the existing user
type UpdateRequest struct {
	AccessToken string            `json:"access_token"`
	UpdateToken string            `json:"update_token"`
	Password    string            `json:"password"`
	Name        string            `json:"name"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	NewPassword string            `json:"new_password"`
	Birthday    time.Time         `json:"birthday"`
	Domain      string            `json:"domain"`
	Roles       []string          `json:"roles"`
	Address     Address           `json:"address"`
	Profile     map[string]string `json:"profile"`
}

// UpdateResponse - Returns error if update fails
type UpdateResponse struct {
	UpdateToken string `json:"update_token"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Err         string `json:"err,omitempty"`
}

// Update - updates the user profile
func Update(req *UpdateRequest, dns string) (*UpdateResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/update")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeUpdateResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return &response, nil
}

// decodeUpdateResponse decodes the response from the service
func decodeUpdateResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response UpdateResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
