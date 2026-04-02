package mail

import (
	"fmt"

	"github.com/dangLuan01/ets-api/internal/utils"
)

type ProviderType string

const (
	ProviderMailtrap ProviderType = "mailtrap"
	ProviderResent ProviderType = "resent"
)

type ProviderFactory interface {
	CreateProvider(config *MailConfig) (EmailProviderService, error)
}

type MailtrapProviderFactory struct {}

type ResentProviderFactory struct {}

func (f *MailtrapProviderFactory) CreateProvider(config *MailConfig) (EmailProviderService, error)  {
	return NewMailtrapProvider(config)
}

func (f *ResentProviderFactory) CreateProvider(config *MailConfig) (EmailProviderService, error) {
	return NewResentProvider(config)
}

func NewProviderFactory(providerType ProviderType) (ProviderFactory, error) {
	switch providerType {
	case ProviderMailtrap:
		return &MailtrapProviderFactory{}, nil
	case ProviderResent:
		return &ResentProviderFactory{}, nil
	default:
		return nil, utils.NewError(string(utils.ErrCodeInternal),fmt.Sprintf("Unsupported provider type:%s", providerType))
	}
	
}