package runner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hoshigakikisame/kabarin/pkg/providers"
	"github.com/hoshigakikisame/kabarin/pkg/utils"
)

type runner struct {
	options   *Options
	providers *[]providers.Provider
}

func New(options *Options, providers *[]providers.Provider) *runner {
	return &runner{options: options, providers: providers}
}

func (r *runner) runProviders(cb func(provider providers.Provider) error) error {
	for _, provider := range *r.providers {
		if err := cb(provider); err != nil {
			return err
		}
	}
	return nil
}

func (r *runner) Notify() error {

	if utils.HasStdin() {

		var dataQueue []string
		maxChar := int(r.options.CharLimit)

		if r.options.isBulk {
			dataBytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			dataQueue = append(dataQueue, string(dataBytes))

			for _, line := range dataQueue {
				if maxChar != 0 {
					dataQueue = utils.ChunkTextStream(line, maxChar)
				}

				r.runProviders(func(provider providers.Provider) error {
					for _, data := range dataQueue {
						if err := provider.SendText(&data); err != nil {
							return err
						}
					}
					return nil
				})
			}
		} else {
			sc := bufio.NewScanner(os.Stdin)

			for sc.Scan() {
				stream := sc.Text()

				if maxChar != 0 {
					dataQueue = utils.ChunkTextStream(stream, maxChar)
				} else {
					dataQueue = append(dataQueue, stream)
				}

				r.runProviders(func(provider providers.Provider) error {
					for _, data := range dataQueue {
						if err := provider.SendText(&data); err != nil {
							return err
						}
					}
					return nil
				})

			}
		}
	}

	if r.options.File != "" {
		var outFileName string

		if r.options.ChunkSize == 0 {
			data, err := os.ReadFile(r.options.File)
			if err != nil {
				return err
			}
			outFileName = filepath.Base(r.options.File)

			for _, provider := range *r.providers {
				if err := provider.SendFile(&outFileName, &data); err != nil {
					return err
				}
			}
		} else {
			utils.FileSplit(r.options.File, int(r.options.ChunkSize), func(chunk []byte, iteration int) error {
				outFileName := fmt.Sprintf("%s_pt-%d%s", strings.TrimSuffix(filepath.Base(r.options.File), filepath.Ext(r.options.File)), iteration, filepath.Ext(r.options.File))

				r.runProviders(func(provider providers.Provider) error {
					if err := provider.SendFile(&outFileName, &chunk); err != nil {
						return err
					}
					return nil
				})

				return nil
			})
		}
	}

	return nil
}
