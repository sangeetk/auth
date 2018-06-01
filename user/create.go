package user

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Create - Creates a new User into filesytem
func (u *User) Create() error {
	// Create BaseDir if it doesn't exist
	dir := BaseDir(u.Username)
	if err := os.MkdirAll(dir, os.ModeDir|0755); err != nil {
		return err
	}

	// Create file
	path := Path(u.Username)
	userJSON, _ := json.MarshalIndent(u, "", "    ")
	if err := ioutil.WriteFile(path, userJSON, 0644); err != nil {
		return err
	}
	return nil
}
