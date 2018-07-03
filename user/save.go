package user

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Save - Updates the user details
func (u *User) Save() error {
	path := Path(u.Username)
	u.UpdatedAt = time.Now()

	userJSON, _ := json.MarshalIndent(u, "", "    ")
	err := ioutil.WriteFile(path, userJSON, 0644)
	return err
}
