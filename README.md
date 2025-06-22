# Kabarin

Kabarin is a notification utility that take stdin and a file as an input and send it to specified service (currently supported: telegram). Inspired by amazing [projectdiscovery](https://github.com/projectdiscovery) tool [notify](https://github.com/projectdiscovery/notify)

## What Sets Kabarin Apart?
- **MTProto**: Kabarin uses MTProto to send message to telegram, which is faster, more reliable and support more file size than telegram bot API
- **File Processing**: Kabarin can send file as an attachment
- **File Splitter**: Option to split file as specified chunk size

## Installation
```bash
go install -v github.com/hoshigakikisame/kabarin/cmd/kabarin@latest
```

## Usage
```text
$ kabarin -h

   __        __            _
  / /_____  / /  ___  ____(_)__
 /  '_/ _ \/ _ \/ _ \/ __/ / _ \
/_/\_\\_,_/_.__/\_,_/_/ /_/_//_/  v0.0.4

by @ferdirianrk

Usage:
  kabarin <options>
  <input> | kabarin <options>

Options:
  -f,  -file <FILE>               File to be send
  -cl, -char-limit <CHAR_LIMIT>   Characters limit in single request (default: 0 (unlimited))
  -cs, -chunk-size <CHUNK_SIZE>   Size of chunks produced by splitting input file (in MB)
  -b,  -bulk                      Enable bulk processing
  -rl, -rate-limit <RATE_LIMIT>   Maximum notification to send per second (default: 1)
  -d,  -delay <DELAY>             Delay in seconds between each notification
  -v,  -version                   Show kabarin version
```

## Configuration
Kabarin use environment variable for configuration. You can set it in your shell or create `.env` file in the same directory as kabarin binary. Here is the list of available configuration:
- `KABARIN_TELEGRAM_API_ID`: Telegram API ID
- `KABARIN_TELEGRAM_API_HASH`: Telegram API Hash
- `TELEGRAM_BOT_TOKEN`: Telegram Bot Token
- `TELEGRAM_RECEIVER_ID`: Telegram Receiver ID (t.me/[username])

## How to use
### Stream text piped input
```bash
subfinder -d hackerone.com | kabarin
```
![alt text](assets/stream_text.png)

### Bulk text piped input
```bash
subfinder -d hackerone.com | kabarin -bulk
```
![alt text](assets/bulk_text.png)

### Send file as attachment
```bash
echo example.com | assetfinder --subs-only > subs.txt; kabarin -file subs.txt
```
![alt text](assets/file.png)

### Send file as attachment with specified chunk size in MB
```bash
echo example.com | assetfinder --subs-only > subs.txt; kabarin -file subs.txt -cs 10
```
![alt text](assets/chunk_file.png)

## TROUBLESHOOT
- `go install` might just froze and if you get `Out of memory` too from `dmesg` after this affair, try to append `GOMEMLIMIT=1GiB` in front of the install command. This issue is typical in low end VPS and it is evidently caused by `gotd/td` excessive ram consumsion during compilation.

## TODO
- [ ] Create unit test
- [x] Add file splitter if input file size is too big
- [x] Add bulk message processing support
- [x] Add max character limit option for text message
- [x] Add request / second limit
- [ ] Add more service (discord, slack, etc)
- [ ] Add more configuration (e.g. retry, timeout, etc)