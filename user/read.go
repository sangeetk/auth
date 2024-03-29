package user

import (
	"encoding/json"
	"io/ioutil"
)

// Read - Reads the user information from file
func Read(u string) (*User, error) {
	var user = new(User)

	// Read file
	path := Path(u)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, user)
	return user, err
}

// Exists returns true of user is already registered else false
func Exists(u string) bool {
	user, _ := Read(u)
	return user != nil
}
