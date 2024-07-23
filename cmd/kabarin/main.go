package main

import (
	"github.com/hoshigakikisame/kabarin/internal/runner"
	"github.com/hoshigakikisame/kabarin/pkg/providers"
	"github.com/hoshigakikisame/kabarin/pkg/providers/telegram"
)

func main() {
	var (
		options   *runner.Options       = runner.Parse()
		providers *[]providers.Provider = &[]providers.Provider{}
	)

	telegramProvider, err := telegram.New()
	if err != nil {
		panic(err)
	}

	*providers = append(*providers, telegramProvider)

	kabarinRunner := runner.New(options, providers)
	if err := kabarinRunner.Notify(); err != nil {
		panic(err)
	}

}
