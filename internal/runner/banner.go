package runner

import (
	"github.com/projectdiscovery/gologger"
)

const (
	author = "ferdirianrk"
	banner = `                           
   __        __            _    
  / /_____  / /  ___  ____(_)__ 
 /  '_/ _ \/ _ \/ _ \/ __/ / _ \
/_/\_\\_,_/_.__/\_,_/_/ /_/_//_/`
)

func showBanner() {
	gologger.Print().Msgf(`
%s  v%s

by @%s
`, banner, version, author)
}
