package utils

import (
	"errors"
	"fmt"
	"os"
)

func HasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	mode := stat.Mode()

	isPipedFromChrDev := (mode & os.ModeCharDevice) == 0
	isPipedFromFIFO := (mode & os.ModeNamedPipe) != 0

	return isPipedFromChrDev || isPipedFromFIFO
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func ValidateEnvVars(keys ...string) error {
	var errStr string
	for _, key := range keys {
		if _, exist := os.LookupEnv(key); !exist {
			errStr += fmt.Sprintf("ENV Var: %s is required\n", key)
		}
	}
	return errors.New(errStr)
}
