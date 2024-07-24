package runner

import (
	"fmt"
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

	if r.options.Data != "" {
		r.runProviders(func(provider providers.Provider) error {
			if err := provider.SendText(&r.options.Data); err != nil {
				return err
			}
			return nil
		})
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

				for _, provider := range *r.providers {
					if err := provider.SendFile(&outFileName, &chunk); err != nil {
						return err
					}
				}

				return nil
			})
		}
	}

	return nil
}
