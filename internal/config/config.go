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

// type Oauth2GoogleConfig struct {
// 	ClientID string
// 	ClientSecret string
// 	RedirectURL string
// 	Scopes []string
// 	Endpoint oauth2.Endpoint
// }

type Config struct {
	ServerAddress string
	DB DatabaseConfig
	MailProviderType string
	MailProviderConfig map[string]any
	Oauth2 *oauth2.Config
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

	oauth2Config := &oauth2.Config{
		ClientID: utils.GetEnv("GOOGLE_CLIENT_ID",""),
		ClientSecret: utils.GetEnv("GOOGLE_CLIENT_SECRET",""),
		RedirectURL: utils.GetEnv("REDIRECT_URL", ""),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
	
	return &Config{
		ServerAddress: fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080")),
		DB: DatabaseConfig {
			Host: utils.GetEnv("DB_HOST",""),
			Port: utils.GetEnv("DB_PORT",""),
			User: utils.GetEnv("DB_USER",""),
			Password: utils.GetEnv("DB_PASSWORD",""),
			DBName: utils.GetEnv("DB_DBNAME",""),
			SSLMode: utils.GetEnv("DB_SSLMODE","disable"),
		},
		MailProviderType: mailProviderType,
		MailProviderConfig: mailProviderConfig,
		Oauth2: oauth2Config,
	}
}

func (c *Config) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
    	c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName,
	)
}