package utils

import (
	"strings"
)

func CheckUserPass(username, password string) bool {
	val := "a"

	if val == username {
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
