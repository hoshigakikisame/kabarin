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
	File      string
	ChunkSize uint
	Version   bool
	Data      string
}

var o *Options

func init() {
	o = &Options{}

	flag.StringVar(&o.File, "f", "", "")
	flag.StringVar(&o.File, "file", "", "")

	flag.UintVar(&o.ChunkSize, "s", 0, "")
	flag.UintVar(&o.ChunkSize, "size", 0, "")

	flag.BoolVar(&o.Version, "v", false, "")
	flag.BoolVar(&o.Version, "version", false, "")

	flag.Usage = func() {
		showBanner()
		help := fmt.Sprintf(`
Usage: %s

Options: %s
`, usage, options)
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
		return nil
	}

	if o.File != "" {
		if !utils.FileExists(o.File) {
			return errors.New("file doesn't exist")
		}

		if o.ChunkSize != 0 {
			o.ChunkSize *= 1024 * 1024
		}

		return nil
	}

	return nil
}

func Parse() *Options {
	showBanner()

	if err := o.validate(); err != nil {
		panic(err)
	}

	return o
}
