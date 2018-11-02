package user

import (
	"crypto/md5"
	"fmt"
)

// DBPath - Basedirectory for storing user profiles
var DBPath string

// Hash - Gets md5 hash of a given string
func Hash(u string) string {
	hash := md5.Sum([]byte(u))
	return fmt.Sprintf("%x", hash)
}

// BaseDir - Creates base directory for user
func BaseDir(u string) string {
	hash := Hash(u)
	dir := DBPath

	for i := 0; i < 10; i += 2 {
		dir += "/" + fmt.Sprintf("%c%c", hash[i], hash[i+1])
	}
	return dir
}

// Path - The full path to user's profile
func Path(u string) string {
	hash := Hash(u)
	path := BaseDir(u)
	path += "/" + hash + ".json"
	return path
}
