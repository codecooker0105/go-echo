package utils

import "os"

func RootDir() string {
	dir, _ := os.Getwd()
	return dir
}
