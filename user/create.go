package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

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
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	userJSON, _ := json.MarshalIndent(u, "", "    ")
	if err := ioutil.WriteFile(path, userJSON, 0644); err != nil {
		return err
	}
	return nil
}
