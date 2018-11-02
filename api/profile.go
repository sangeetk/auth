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

// ProfileRequest - Request the profile for the logged-in user
type ProfileRequest struct {
	AccessToken string `json:"access_token"`
}

// ProfileResponse - Returns the profile for the logged-in user
type ProfileResponse struct {
	Username      string              `json:"username"`
	Name          string              `json:"name"`
	FirstName     string              `json:"first_name"`
	LastName      string              `json:"last_name"`
	Email         string              `json:"email"`
	Birthday      time.Time           `json:"birthday"`
	InitialDomain string              `json:"initial_domain"`
	Roles         map[string][]string `json:"roles"`
	Address       Address             `json:"address"`
	Profile       UserProfile         `json:"profile"`
	Err           string              `json:"err,omitempty"`
}

// Profile - returns the full profile of the user
func Profile(req *ProfileRequest, dns string) (*ProfileResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/profile")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeProfileResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ProfileResponse)
	return &response, nil
}

// decodeProfileResponse decodes the response from the service
func decodeProfileResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response ProfileResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
