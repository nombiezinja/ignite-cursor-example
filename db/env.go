package db

import (
	"os"
)

func Setenv() {
	os.Setenv("address", "127.0.0.1:12345")
	os.Setenv("cache", "default")
}
