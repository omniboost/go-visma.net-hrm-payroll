package vismanet

import (
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const ()

type Oauth2Config struct {
	clientcredentials.Config
}

func NewOauth2Config(tenantID string) *Oauth2Config {
	config := &Oauth2Config{
		Config: clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{},
			TokenURL:     "https://connect.visma.com/connect/token",
			AuthStyle:    oauth2.AuthStyleInParams,
			EndpointParams: url.Values{
				"tenant_id": []string{tenantID},
			},
		},
	}

	return config
}
