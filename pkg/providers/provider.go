package providers

type Provider interface {
	SendText(text *string) error
	SendFile(fileName *string, data *[]byte) error
	Close() error
}
