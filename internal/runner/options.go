package runner

import (
	"errors"
	"flag"

	"github.com/hoshigakikisame/kabarin/pkg/utils"
	"github.com/projectdiscovery/gologger"
)

type Options struct {
	File      string
	Provider  string
	ChunkSize uint
	Version   bool
	CharLimit uint
	isBulk    bool
	RateLimit uint
	Delay     uint
}

var o *Options

func init() {
	o = &Options{}

	flag.StringVar(&o.File, "f", "", "")
	flag.StringVar(&o.File, "file", "", "")

	flag.StringVar(&o.Provider, "p", "", "")
	flag.StringVar(&o.Provider, "Provider", "", "")

	flag.UintVar(&o.ChunkSize, "cs", 0, "")
	flag.UintVar(&o.ChunkSize, "chunk-size", 0, "")

	flag.UintVar(&o.CharLimit, "cl", 0, "")
	flag.UintVar(&o.CharLimit, "char-limit", 0, "")

	flag.BoolVar(&o.isBulk, "b", false, "")
	flag.BoolVar(&o.isBulk, "bulk", false, "")

	flag.UintVar(&o.RateLimit, "rl", 1, "")
	flag.UintVar(&o.RateLimit, "rate-limit", 1, "")

	flag.UintVar(&o.Delay, "d", 0, "")
	flag.UintVar(&o.Delay, "delay", 0, "")

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

	if o.Version {
		showVersion()
	}

	showBanner()

	if err := o.validate(); err != nil {
		gologger.Fatal().Msg(err.Error())
	}

	return o
}
