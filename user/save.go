package user

import (
	"io/ioutil"
	"os"
)

func (u *User) Save() error {
	path := Path(u.Username)

	content := ""

	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	return err
}
