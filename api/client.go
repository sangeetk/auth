package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AuthClient structure
type AuthClient struct {
	DNS string
}

// NewAuthClient - create a new client
func (a *AuthClient) NewAuthClient(dns string) {
	a.DNS = dns
}

// encodeRequest encodes the request as JSON
func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeResponse decodes the response from the service
func decodeResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var response interface{}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
