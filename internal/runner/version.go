package runner

import (
	"os"

	"github.com/projectdiscovery/gologger"
)

const version = "0.0.1"

func showVersion() {
	gologger.Print().Msgf("kabarin v%s", version)
	os.Exit(2)
}
