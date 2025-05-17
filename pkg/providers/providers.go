package providers

import (
	"github.com/hoshigakikisame/kabarin/pkg/providers/telegram"
	"github.com/hoshigakikisame/kabarin/pkg/utils/throttle"
	"github.com/projectdiscovery/gologger"
)

type Providers struct {
	providerList *[]Provider
	throttle     *throttle.Throttle
}

func New(rateLimit int, delay int) (*Providers, error) {
	telegramProvider, err := telegram.New()
	if err != nil {
		return nil, err
	}

	throttle, err := throttle.New(rateLimit, delay)
	if err != nil {
		return nil, err
	}

	throttle.Run()

	return &Providers{
		providerList: &[]Provider{
			telegramProvider,
		},
		throttle: throttle,
	}, nil
}

func (p *Providers) SendText(text *string, delay *uint) error {
	for _, provider := range *p.providerList {
		p.throttle.AddJob(func() {
			provider.SendText(text)
		})
	}
	p.throttle.Wait()
	gologger.Print().Msg(*text)
	return nil
}

func (p *Providers) SendFile(fileName *string, data *[]byte, delay *uint) error {
	for _, provider := range *p.providerList {
		p.throttle.AddJob(func() {
			provider.SendFile(fileName, data)
		})
	}
	p.throttle.Wait()
	gologger.Print().Msg(*fileName)
	return nil
}

func (p *Providers) Close() error {
	for _, provider := range *p.providerList {
		if err := provider.Close(); err != nil {
			return err
		}
	}
	return nil
}