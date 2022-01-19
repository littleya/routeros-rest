package routeros

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
	url    string
}

func NewClient(address, username, password string, tlsConfig *tls.Config) (*Client, error) {
	if !strings.HasPrefix(address, "https://") {
		address = fmt.Sprintf("https://%s", address)
	}
	if !strings.HasSuffix(address, "/rest") {
		address += "/rest"
	}

	c := resty.New()
	c.SetTLSClientConfig(tlsConfig)
	c.SetBasicAuth(username, password)

	client := &Client{
		client: c,
		url:    address,
	}

	err := client.Verify()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Verify() error {
	_, err := c.GetSystemRouterboard()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed verify client, error: %s", err))
	}
	return nil
}
