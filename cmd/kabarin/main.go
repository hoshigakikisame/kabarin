package main

import (
	"github.com/hoshigakikisame/kabarin/internal/runner"
	"github.com/hoshigakikisame/kabarin/pkg/providers"
)

func main() {
	var options *runner.Options = runner.Parse()

	providers, err := providers.New(1)
	if err != nil {
		panic(err)
	}

	kabarinRunner := runner.New(options, providers)

	if err := kabarinRunner.Notify(); err != nil {
		panic(err)
	}
}
