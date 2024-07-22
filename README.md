# KABARIN

Kabarin is a notification utility that take stdin and file as an input and send it to specified service (currently supported: telegram). Inspired by amazing [projectdiscovery](https://github.com/projectdiscovery) tool [notify](https://github.com/projectdiscovery/notify)

## What Sets Kabarin Apart?
- **MTProto**: Kabarin use MTProto to send message to telegram, which is faster, more reliable and support more file size than telegram bot API

## TODO
- [ ] Create unit test
- [ ] Add bulk message processing support
- [ ] Add request / second limit
- [ ] Add more service (discord, slack, etc)
- [ ] Add more configuration (e.g. retry, timeout, etc)