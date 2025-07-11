package main

import (
	"github.com/hoshigakikisame/kabarin/internal/runner"
	"github.com/hoshigakikisame/kabarin/pkg/providers"
	"github.com/projectdiscovery/gologger"
)

func main() {
	var options *runner.Options = runner.Parse()

	providers, err := providers.New(int(options.RateLimit), int(options.Delay), options.Provider)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}

	kabarinRunner := runner.New(options, providers)

	if err := kabarinRunner.Notify(); err != nil {
		gologger.Fatal().Msg(err.Error())
	}
}
