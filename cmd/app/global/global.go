package global

import "os"

var Secret = []byte(os.Getenv("SECRET_KEY"))

const Userkey = "user"
