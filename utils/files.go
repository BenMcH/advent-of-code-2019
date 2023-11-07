package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) string {
	file, err := os.Open(path)

	if err != nil {
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr []string

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	return strings.Trim(strings.Join(arr, "\n"), "\n")
}
