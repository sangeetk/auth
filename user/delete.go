package user

import (
	"os"
)

func (u *User) Delete() error {
	path := Path(u.Username)
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
