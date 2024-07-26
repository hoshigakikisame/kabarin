package providers

import (
	"time"

	"github.com/hoshigakikisame/kabarin/pkg/providers/telegram"
	"github.com/hoshigakikisame/kabarin/pkg/utils/throttle"
)

type Providers struct {
	providerList *[]Provider
	throttle     *throttle.Throttle
}

func New(rateLimit int) (*Providers, error) {
	telegramProvider, err := telegram.New()
	if err != nil {
		return nil, err
	}

	throttle, err := throttle.New(time.Duration(rateLimit))
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
		wait(*delay)
	}
	p.throttle.Wait()
	return nil
}

func (p *Providers) SendFile(fileName *string, data *[]byte, delay *uint) error {
	for _, provider := range *p.providerList {
		p.throttle.AddJob(func() {
			provider.SendFile(fileName, data)
		})
		wait(*delay)
	}
	p.throttle.Wait()
	return nil
}

func wait(sec uint) {
	if sec > 0 {
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
