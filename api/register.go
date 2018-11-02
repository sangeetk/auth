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

// Address - address fields
type Address struct {
	AddressType string `json:"address_type"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Zip         string `json:"zip"`
}

// UserProfile - used for storing profile infomation
type UserProfile map[string]string

// RegisterRequest - New User registration request
type RegisterRequest struct {
	Username  string      `json:"username"`
	Name      string      `json:"name"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Birthday  time.Time   `json:"birthday"`
	Domain    string      `json:"domain"`
	Roles     []string    `json:"roles"`
	Address   Address     `json:"address"`
	Profile   UserProfile `json:"profile"`
}

// RegisterResponse - New User registration response
type RegisterResponse struct {
	ConfirmToken string `json:"confirm_token,omitempty"`
	UpdateToken  string `json:"update_token,omitempty"`
	Err          string `json:"err,omitempty"`
}

// Register - registers a new user
func Register(req *RegisterRequest, dns string) (*RegisterResponse, error) {
	ctx := context.Background()
	tgt, err := url.Parse("http://" + dns + "/register")
	if err != nil {
		log.Fatal(err.Error())
	}
	endPoint := ht.NewClient("POST", tgt, encodeRequest, decodeRegisterResponse).Endpoint()
	resp, err := endPoint(ctx, req)
	if err != nil {
		return nil, err
	}
	response := resp.(RegisterResponse)
	return &response, nil
}

// decodeRegisterResponse decodes the response from the service
func decodeRegisterResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response RegisterResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
