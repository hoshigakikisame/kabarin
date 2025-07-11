package providers

import (
	"errors"

	"github.com/hoshigakikisame/kabarin/pkg/providers/discord"
	"github.com/hoshigakikisame/kabarin/pkg/providers/telegram"
	"github.com/hoshigakikisame/kabarin/pkg/utils/throttle"
	"github.com/projectdiscovery/gologger"
)

type Providers struct {
	providerList *[]Provider
	throttle     *throttle.Throttle
}

func New(rateLimit int, delay int, provider string) (*Providers, error) {
	var providerList []Provider

	// Define a map of provider names to their respective factory functions
	providerConfigs := map[string][]func() (Provider, error){
		"all": {
			func() (Provider, error) {
				return telegram.New()
			},
			func() (Provider, error) {
				return discord.New()
			},
		},
		"discord": {
			func() (Provider, error) {
				return discord.New()
			},
		},
		"telegram": {
			func() (Provider, error) {
				return telegram.New()
			},
		},
		"": {
			func() (Provider, error) {
				return telegram.New()
			},
		},
	}

	// Check if the specified provider exists in the map
	factories := providerConfigs[provider]
	if factories == nil {
		return nil, errors.New("unknown provider specified: " + provider)
	}

	// Iterate over the factory functions and create instances of the providers
	for _, factory := range factories {
		newProvider, err := factory()
		if err != nil {
			return nil, err
		}
		providerList = append(providerList, newProvider)
	}

	throttle, err := throttle.New(rateLimit, delay)
	if err != nil {
		return nil, err
	}

	throttle.Run()

	return &Providers{
		providerList: &providerList,
		throttle:     throttle,
	}, nil
}

func (p *Providers) SendText(text *string, delay *uint) error {
	for _, provider := range *p.providerList {
		currentProvider := provider
		p.throttle.AddJob(func() {
			currentProvider.SendText(text)
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
