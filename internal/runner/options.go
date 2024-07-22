package runner

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hoshigakikisame/kabarin/pkg/utils"
)

type Options struct {
	File    string
	Version bool
	Data    string
}

var o *Options

func init() {
	o = &Options{}

	flag.StringVar(&o.File, "f", "", "")
	flag.StringVar(&o.File, "file", "", "")

	flag.BoolVar(&o.Version, "v", false, "")
	flag.BoolVar(&o.Version, "version", false, "")

	flag.Usage = func() {
		help := fmt.Sprintf(`
%s  v%s

by @%s

Usage: %s

Options: %s
`, banner, version, author, usage, options)
		fmt.Print(help)
	}

	flag.Parse()
}

func (o *Options) validate() error {
	if utils.HasStdin() {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		o.Data = string(data)
	} else if o.File == "" && !utils.FileExists(o.File) {
		return errors.New("file doesn't exist")
	}
	return nil
}

func Parse() *Options {
	if err := o.validate(); err != nil {
		panic(err)
	}
	return o
}
