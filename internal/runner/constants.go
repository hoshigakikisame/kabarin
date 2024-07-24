package runner

const (
	version = "1.0.0"
	author  = "ferdirianrk"
	banner  = `
    __         __               _     
   / /______ _/ /_  ____  _____(_)___ 
  / //_/ __ \/ __ \/ __ \/ ___/ / __ \
 / ,< / /_/ / /_/ / /_/ / /  / / / / /
/_/|_|\__,_/_.___/\__,_/_/  /_/_/ /_/`
	usage = `
  kabarin <options>
  <input> | kabarin <options>`
	options = `
  -f, -file <FILE>        File to be send
  -s, -size <CHUNK_SIZE>  Size of chunks produced by splitting input file (in MB)
  -v, -version            Show kabarin version`
)
