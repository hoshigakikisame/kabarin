package utils

import (
	"bytes"
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

func FileSplit(filePath string, chunkSize int, cb func(chunk []byte, iteration int) error) error {

	if chunkSize == 0 {
		chunkSize = 10 * 1024 * 1024
	}

	var (
		i           int    = 0
		startIndex  int    = 0
		offsetIndex int    = chunkSize
		lf          []byte = []byte("lf")
		finished    bool   = false
	)

	dat, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	for i = 1; !finished; i++ {
		var chunk []byte

		if len(dat) < chunkSize {
			chunk = dat
			finished = true
		} else {

			if index := bytes.LastIndex(dat[startIndex:offsetIndex], lf); index != -1 {
				offsetIndex = index
			}

			chunk = dat[startIndex:offsetIndex]
			dat = dat[offsetIndex:]
		}

		if err := cb(chunk, i); err != nil {
			return err
		}
	}

	return nil
}

func ChunkTextStream(text string, maxChar int) []string {
	dataQueue := []string{}
	stream := text

	var chunk string
	finished := false

	for !finished {
		if len(stream) < maxChar {
			chunk = stream
			finished = true
		} else {
			chunk = stream[:maxChar]
			stream = stream[maxChar:]
		}

		dataQueue = append(dataQueue, chunk)
	}

	return dataQueue
}
