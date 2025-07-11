package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/projectdiscovery/gologger"
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

	if errStr == "" {
		return nil
	}

	return errors.New(errStr)
}

func TextChunkStream(text string, maxChar uint) <-chan string {

	textQueue := make(chan string)
	var chunk string

	go func() {
		defer close(textQueue)

		for last := false; !last; {
			if len(text) < int(maxChar) || maxChar == 0 {
				chunk = text
				last = true
			} else {
				chunk = text[:maxChar]
				text = text[maxChar:]
			}
			textQueue <- chunk
		}
	}()
	return textQueue
}

func FileChunkStream(filePath string, chunkSize uint) <-chan []byte {

	if chunkSize == 0 {
		chunkSize = 10 * 1024 * 1024
	}

	file, err := os.Open(filePath)
	if err != nil {
		gologger.Fatal().Msgf("Error on opening file, detail: %v", err)
	}

	chunkQueue := make(chan []byte)
	chunk := make([]byte, chunkSize)

	go func() {
		defer close(chunkQueue)

		for last := false; !last; {

			n, err := file.Read(chunk)

			if err != nil {
				gologger.Fatal().Msgf("Error during file read, detail: %v", err)
			}

			if n < int(chunkSize) {
				chunk = chunk[:n]
				last = true
			}

			chunkQueue <- chunk
		}

	}()

	return chunkQueue
}
