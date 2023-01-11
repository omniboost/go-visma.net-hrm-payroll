package vismanet_test

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	vismanet "github.com/omniboost/go-visma.net-hrm-payroll"
)

var (
	client *vismanet.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tenantID := os.Getenv("TENANT_ID")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := vismanet.NewOauth2Config(tenantID)
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client = vismanet.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
