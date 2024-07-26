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
```bash
kabarin -h
```

## Configuration
Kabarin use environment variable for configuration. You can set it in your shell or create `.env` file in the same directory as kabarin binary. Here is the list of available configuration:
- `KABARIN_TELEGRAM_API_ID`: Telegram API ID
- `KABARIN_TELEGRAM_API_HASH`: Telegram API Hash
- `TELEGRAM_BOT_TOKEN`: Telegram Bot Token
- `TELEGRAM_RECEIVER_ID`: Telegram Receiver ID (t.me/[username])

## Usage
### Stream text piped input
```bash
echo hackerone.com | assetfinder --subs-only | kabarin
```
![alt text](https://placehold.co/400x600)

### Bulk text piped input
```bash
echo hackerone.com | assetfinder --subs-only | kabarin -bulk
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

## TODO
- [ ] Create unit test
- [x] Add file splitter if input file size is too big
- [x] Add bulk message processing support
- [x] Add max character limit option for text message
- [x] Add request / second limit
- [ ] Add more service (discord, slack, etc)
- [ ] Add more configuration (e.g. retry, timeout, etc)