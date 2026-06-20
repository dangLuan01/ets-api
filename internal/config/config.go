package config

import (
	"fmt"

	"github.com/dangLuan01/ets-api/internal/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"	
)

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}

type Config struct {
	ServerAddress string
	DB DatabaseConfig
	MailProviderType string
	MailProviderConfig map[string]any
	OAuth2ProviderConfig map[string]*oauth2.Config

	//Oauth2 *oauth2.Config
}

func NewConfig() *Config {

	mailProviderConfig := make(map[string]any)
	
	mailProviderType := utils.GetEnv("MAIL_PROVIDER_TYPE","mailtrap")
	if mailProviderType == "mailtrap" {
		mailtrapConfig := map[string]any {
			"mail_sender": utils.GetEnv("MAILTRAP_MAIL_SENDER","no-reply@test.com"),
			"name_sender": utils.GetEnv("MAILTRAP_NAME_SENDER","Support Team"),
			"mailtrap_url": utils.GetEnv("MAILTRAP_URL",""),
			"mailtrap_api_key": utils.GetEnv("MAILTRAP_API_KEY",""),
		}

		mailProviderConfig["mailtrap"] = mailtrapConfig
	}

	if mailProviderType == "resent" {
		resentConfig := map[string]any{
			"mail_sender": utils.GetEnv("RESENT_MAIL_SENDER","no-reply@app.xoailac.top"),
			"name_sender": utils.GetEnv("RESENT_NAME_SENDER","Support Team"),
			"resent_api_key": utils.GetEnv("RESENT_API_KEY",""),
		}

		mailProviderConfig["resent"] = resentConfig
	}

	oauth2ProviderConfig := make(map[string]*oauth2.Config)

	oauth2ProviderConfig["google"] = &oauth2.Config{
		ClientID: utils.GetEnv("GOOGLE_CLIENT_ID",""),
		ClientSecret: utils.GetEnv("GOOGLE_CLIENT_SECRET",""),
		RedirectURL: utils.GetEnv("GOOGLE_REDIRECT_URL", ""),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}

	oauth2ProviderConfig["facebook"] = &oauth2.Config{
		ClientID: utils.GetEnv("FACEBOOK_CLIENT_ID",""),
		ClientSecret: utils.GetEnv("FACEBOOK_CLIENT_SECRET",""),
		RedirectURL: utils.GetEnv("FACEBOOK_REDIRECT_URL", ""),
		Scopes: []string{"public_profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v3.2/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v3.2/oauth/access_token",
		},
	}
	
	return &Config{
		ServerAddress: fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")),
		DB: DatabaseConfig {
			Host: utils.GetEnv("DB_HOST",""),
			Port: utils.GetEnv("DB_PORT",""),
			User: utils.GetEnv("DB_USER",""),
			Password: utils.GetEnv("DB_PASSWORD",""),
			DBName: utils.GetEnv("DB_DBNAME",""),
			SSLMode: utils.GetEnv("DB_SSLMODE","false"),
		},
		MailProviderType: mailProviderType,
		MailProviderConfig: mailProviderConfig,
		OAuth2ProviderConfig: oauth2ProviderConfig,
	}
}

func (c *Config) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s",
    	c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName, c.DB.SSLMode,
	)
}