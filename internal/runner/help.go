package runner

import "github.com/projectdiscovery/gologger"

const (
	usage = `
  kabarin <options>
  <input> | kabarin <options>`
	options = `
  -f, -file <FILE>        File to be send
  -cl, -char-limit <CHAR_LIMIT>  Characters limit in single request [default: 0 (unlimited)]
  -cs, -chunk-size <CHUNK_SIZE>  Size of chunks produced by splitting input file (in MB)
  -v, -version            Show kabarin version`
)

func showUsage() {
	gologger.Print().Msgf(usage)
}

func showOptions() {
	gologger.Print().Msgf(options)
}
