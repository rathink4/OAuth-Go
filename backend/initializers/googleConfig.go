package initializers

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthConfig struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig OAuthConfig

func GoogleConfig() oauth2.Config {
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/oauth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig
}
