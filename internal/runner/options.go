package runner

import (
	"errors"
	"flag"

	"github.com/hoshigakikisame/kabarin/pkg/utils"
)

type Options struct {
	File      string
	ChunkSize uint
	Version   bool
	CharLimit uint
	isBulk    bool
}

var o *Options

func init() {
	o = &Options{}

	flag.StringVar(&o.File, "f", "", "")
	flag.StringVar(&o.File, "file", "", "")

	flag.UintVar(&o.ChunkSize, "cs", 0, "")
	flag.UintVar(&o.ChunkSize, "chunk-size", 0, "")

	flag.UintVar(&o.CharLimit, "cl", 0, "")
	flag.UintVar(&o.CharLimit, "char-limit", 0, "")

	flag.BoolVar(&o.isBulk, "b", false, "")
	flag.BoolVar(&o.isBulk, "bulk", false, "")

	flag.BoolVar(&o.Version, "v", false, "")
	flag.BoolVar(&o.Version, "version", false, "")

	flag.Usage = func() {
		showBanner()
		showUsage()
		showOptions()
	}
	flag.Parse()
}

func (o *Options) validate() error {

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
