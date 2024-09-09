package mockoidc

import (
	"crypto/tls"
	"github.com/jsiebens/mockoidc/pkg/gen/mockoidc/v1/mockoidcv1connect"
	"net/http"
)

func NewClient(serverURL string, insecureSkipVerify bool) mockoidcv1connect.MockOIDCServiceClient {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return mockoidcv1connect.NewMockOIDCServiceClient(client, serverURL)
}
