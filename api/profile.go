package api

import (
	"context"
	"log"
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
	Profile       Profile             `json:"profile"`
	Err           string              `json:"err,omitempty"`
}

// Profile - returns the full profile of the user
func (a *AuthClient) Profile(req *ProfileRequest) (*ProfileResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + a.DNS + "/profile")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(ProfileResponse)
	return &response, nil
}
