package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/auth/token"
	"git.urantiatech.com/auth/auth/user"
	"github.com/urantiatech/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

// Register - Register a new User
func (Auth) Register(ctx context.Context, req api.RegisterRequest) (api.RegisterResponse, error) {
	var response = api.RegisterResponse{}

	var passwordHash []byte
	var err error

	if req.CacheKey == "" {
		if req.Email == "" || req.Password == "" {
			response.Err = ErrorInvalidRequest.Error()
			return response, nil
		}

		// Use email as username if empty
		if req.Username == "" {
			req.Username = req.Email
		}

		passwordHash, err = bcrypt.GenerateFromPassword([]byte(req.Password), 11)
		if err != nil {
			log.Println("Bcrypt error:", err.Error())
		}
	}

	var u = &user.User{
		Username:      req.Username,
		Name:          req.Name,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Password:      passwordHash,
		Birthday:      req.Birthday,
		InitialDomain: req.Domain,
		Confirmed:     false,
	}

	u.Roles = make(map[string][]string)
	u.Roles[req.Domain] = req.Roles

	u.Address = req.Address
	u.Profile = req.Profile

	// Multi-step registration
	if req.CacheReq || req.CacheKey != "" {
		// Step 1
		if req.CacheReq && req.CacheKey == "" {
			// Check if user already exists
			if user.Exists(u.Username) {
				response.Err = ErrorAlreadyRegistered.Error()
				return response, nil
			}
			// Create a new token
			cacheKey, err := token.NewToken(u, req.Domain, user.TemporaryRegistrationValidity)
			if err != nil {
				response.Err = ErrorInvalidRequest.Error()
				return response, nil
			}
			// Insert user into the Cache
			err = user.TemporaryRegistration.Add(cacheKey, u, user.TemporaryRegistrationValidity)
			if err != nil {
				response.Err = ErrorInvalidRequest.Error()
				return response, nil
			}

			response.CacheKey = cacheKey
			return response, nil
		}

		// Step 2
		if req.CacheReq && req.CacheKey != "" {
			// Get the user from the Cache
			cachedUser, ok := user.TemporaryRegistration.Get(req.CacheKey)
			if !ok {
				response.Err = ErrorInvalidRequest.Error()
				return response, nil
			}

			// Update the user
			updateCachedUser(cachedUser.(*user.User), &req)

			// Create a new token
			cacheKey, err := token.NewToken(u, req.Domain, user.TemporaryRegistrationValidity)
			if err != nil {
				response.Err = ErrorInvalidRequest.Error()
				return response, nil
			}

			// Update Cache
			u = cachedUser.(*user.User)
			err = user.TemporaryRegistration.Add(cacheKey, u, user.TemporaryRegistrationValidity)
			if err != nil {
				response.Err = ErrorInvalidRequest.Error()
				return response, nil
			}
			response.CacheKey = cacheKey
			return response, nil
		}

		// Step 3
		if !req.CacheReq && req.CacheKey != "" {
			// Get the user from the Cache
			cachedUser, ok := user.TemporaryRegistration.Get(req.CacheKey)
			if !ok {
				response.Err = ErrorInvalidRequest.Error()
				return response, nil
			}

			// Update the user
			updateCachedUser(cachedUser.(*user.User), &req)

			// Delete the user from Cache by setting expiry to 1 second
			user.TemporaryRegistration.Set(req.CacheKey, nil, time.Second)

			// Create the user
			u = cachedUser.(*user.User)
		}

	}

	// Register User
	if err := u.Create(); err != nil {
		response.Err = ErrorAlreadyRegistered.Error()
		return response, nil
	}

	// Create the Confirmation token
	response.ConfirmToken, err = token.NewToken(u, req.Domain, token.ConfirmTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	// Create the Update token
	response.UpdateToken, err = token.NewToken(u, req.Domain, token.UpdateTokenValidity)
	if err != nil {
		response.Err = err.Error()
		return response, nil
	}

	response.Username = u.Username
	response.FirstName = u.FirstName
	response.LastName = u.LastName
	response.Email = u.Email
	response.Domain = req.Domain
	response.Roles = u.GetRoles(req.Domain)

	return response, nil
}

func updateCachedUser(u *user.User, req *api.RegisterRequest) {
	// Update user fields
	if req.Name != "" {
		u.Name = req.Name
	}
	if req.FirstName != "" {
		u.FirstName = req.FirstName
	}
	if req.LastName != "" {
		u.LastName = req.LastName
	}
	/*
		if req.Email != "" {
			u.Email = req.Email
		}
	*/
	if !req.Birthday.IsZero() {
		u.Birthday = req.Birthday
	}

	// Address information
	if req.Address.AddressType != "" {
		u.Address.AddressType = req.Address.AddressType
	}
	if req.Address.Address1 != "" {
		u.Address.Address1 = req.Address.Address1
	}
	if req.Address.Address2 != "" {
		u.Address.Address2 = req.Address.Address2
	}
	if req.Address.City != "" {
		u.Address.City = req.Address.City
	}
	if req.Address.State != "" {
		u.Address.State = req.Address.State
	}
	if req.Address.Country != "" {
		u.Address.Country = req.Address.Country
	}
	if req.Address.Zip != "" {
		u.Address.Zip = req.Address.Zip
	}

	/*
		// Update Roles
		if len(req.Roles) > 0 {
			u.Roles[t.Domain] = req.Roles
		}
	*/

	// Profile information
	for k, v := range req.Profile {
		if u.Profile == nil {
			u.Profile = make(map[string]string)
		}
		if v == "" {
			delete(u.Profile, k)
		} else {
			u.Profile[k] = v
		}
	}
}

// MakeRegisterEndpoint -
func MakeRegisterEndpoint(svc Auth) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.RegisterRequest)
		return svc.Register(ctx, req)
	}
}

// DecodeRegisterRequest -
func DecodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
