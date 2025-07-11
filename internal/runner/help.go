package runner

import "github.com/projectdiscovery/gologger"

const (
	usage = `
Usage:
  kabarin <options>
  <input> | kabarin <options>`
	options = `
Options:
  -f,  -file <FILE>               File to be send
  -p,  -provider <PROVIDER>       Target provider to be send (default: telegram (telegram, discord, all))
  -cl, -char-limit <CHAR_LIMIT>   Characters limit in single request (default: 0 (unlimited))
  -cs, -chunk-size <CHUNK_SIZE>   Size of chunks produced by splitting input file (in MB)
  -b,  -bulk                      Enable bulk processing
  -rl, -rate-limit <RATE_LIMIT>   Maximum notification to send per second (default: 1)
  -d,  -delay <DELAY>             Delay in seconds between each notification
  -v,  -version                   Show kabarin version
  `
)

func showUsage() {
	gologger.Print().Msgf(usage)
}

func showOptions() {
	gologger.Print().Msgf(options)
}
