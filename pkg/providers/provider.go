package providers

type Provider interface {
	SendText(text string) error
	SendFile(filepath string) error
	Close() error
}
