package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

// TemporaryRegistrationValidity for multi-step registrations
var TemporaryRegistrationValidity time.Duration

// TemporaryRegistration - Cache to store invalid access tokens
var TemporaryRegistration *cache.Cache

// Create - Creates a new User into filesytem
func (u *User) Create() error {
	// Create BaseDir if it doesn't exist
	dir := BaseDir(u.Username)
	if err := os.MkdirAll(dir, os.ModeDir|0755); err != nil {
		return err
	}

	path := Path(u.Username)

	// Check if file already exists
	_, err := ioutil.ReadFile(path)
	if err == nil {
		return errors.New("User Already Exists")
	}

	// Create file
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	u.PasswordUpdatedAt = now

	userJSON, _ := json.MarshalIndent(u, "", "    ")
	if err := ioutil.WriteFile(path, userJSON, 0644); err != nil {
		return err
	}
	return nil
}
