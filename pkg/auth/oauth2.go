package auth

import (
	"context"
	"encoding/json"

	"github.com/dangLuan01/ets-api/internal/config"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"golang.org/x/oauth2"
)

type oauth2Service struct {
	conf config.Config
}

func NewOauth2Service(conf config.Config) Oauth2Service {
	return &oauth2Service{
		conf: conf,
	}
}

func (os *oauth2Service) OAuth2Login(provider string) (string, string, error) {
	state, err := utils.GenerateRandomString(6)
	if err != nil {
		return "", "", utils.NewError(string(utils.ErrCodeInternal), " Vui lòng thử lại sau.")
	}
	
	url := os.conf.OAuth2ProviderConfig[provider].AuthCodeURL(state, oauth2.AccessTypeOnline)

	return url, state, nil
}

func (os *oauth2Service) OAuth2Callback(provider, code string) (models.User, error) {
	var (
		user models.User
		googleUser models.GoogleUser
		facebookUser models.FacebookUser
	)

	tok, err := os.conf.OAuth2ProviderConfig[provider].Exchange(context.Background(), code)
	if err != nil {
		return models.User{}, err
	}

	client := os.conf.OAuth2ProviderConfig[provider].Client(context.Background(), tok)

	switch provider {
	case "facebook":
		resp, err := client.Get("https://graph.facebook.com/me?fields=id,name,email,picture")
			if err != nil {
			return models.User{}, err
		}

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&facebookUser); err != nil {
			return models.User{}, err
		}

		user = models.User{
			UserName: facebookUser.Name,
			Email: facebookUser.Email,
			Avatar: &facebookUser.Picture.Data.Url,
			OpenID: &facebookUser.Id,
		}
		
	default:
		resp, err := client.Get("https://openidconnect.googleapis.com/v1/userinfo")
			if err != nil {
			return models.User{}, err
		}

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
			return models.User{}, err
		}

		user = models.User{
			UserName: googleUser.Name,
			Email: googleUser.Email,
			Avatar: &googleUser.Picture,
			OpenID: &googleUser.Sub,
		}
	}
	
	return user, nil
}