package db

import (
	"os"
)

func Setenv() {
	os.Setenv("address", "127.0.0.1:10800")
	os.Setenv("cache", "default")
}
