package common

import (
	"bufio"
	"strings"
)

func Getenv(key string, scanner *bufio.Scanner) (value string) {
	for scanner.Scan() {
		key = key + "="
		line := scanner.Text()
		if strings.HasPrefix(line, key) {
			value = strings.TrimPrefix(line, key)
			return value
		}
	}
	return
}
