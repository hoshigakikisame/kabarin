package runner

import (
	"github.com/hoshigakikisame/kabarin/pkg/providers"
)

type runner struct {
	options   *Options
	providers *[]providers.Provider
}

func New(options *Options, providers *[]providers.Provider) *runner {
	return &runner{options: options, providers: providers}
}

func (r *runner) Notify() error {
	for _, provider := range *r.providers {
		defer provider.Close()
		if r.options.Data != "" {
			if err := provider.SendText(r.options.Data); err != nil {
				return err
			}
		} else {
			if err := provider.SendFile(r.options.File); err != nil {
				return err
			}
		}
	}
	return nil
}
