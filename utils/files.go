package utils

import (
	"os"
)

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	return string(bytes)
}
