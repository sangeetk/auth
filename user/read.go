package user

import (
	"encoding/json"
	"io/ioutil"
)

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