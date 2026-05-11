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

func (os *oauth2Service) GoogleLogin() (string, error) {
	state, err := utils.GenerateRandomString(6)
	if err != nil {
		return "", utils.NewError(string(utils.ErrCodeInternal), " Vui lòng thử lại sau.")
	}
	
	url := os.conf.Oauth2.AuthCodeURL(state, oauth2.AccessTypeOnline)

	return url, nil
}

func (os *oauth2Service) GoogleCallback(code string) (models.GoogleUser, error) {
	var user models.GoogleUser
	tok, err := os.conf.Oauth2.Exchange(context.Background(), code)
	if err != nil {
		return models.GoogleUser{}, err
	}

	client := os.conf.Oauth2.Client(context.Background(), tok)
	resp, err := client.Get("https://openidconnect.googleapis.com/v1/userinfo")
	if err != nil {
		return models.GoogleUser{}, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return models.GoogleUser{}, err
	}

	return user, nil
}