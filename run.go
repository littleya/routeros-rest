package routeros

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/tidwall/gjson"
)

func (c *Client) GET(url string) (gjson.Result, error) {
	resp, err := c.client.R().Get(url)

	if err != nil {
		return gjson.Result{}, err
	} else if resp.RawResponse.StatusCode/200 != 1 {
		return gjson.Result{}, fmt.Errorf(fmt.Sprintf("Response status code %d, response: %s", resp.RawResponse.StatusCode, resp.Body()))
	}

	respParse := gjson.Parse(string(resp.Body()))

	if respParse.Get("detail").String() != "" {
		return respParse, fmt.Errorf(respParse.Get("detail").String())
	}

	return respParse, nil
}

func (c *Client) GetSystemRouterboard() (SystemRouterboard, error) {
	url := c.generateGetUrl("/system/routerboard", map[string]string{})

	resp, err := c.GET(url)
	if err != nil {
		return SystemRouterboard{}, err
	}

	r := SystemRouterboard{}
	err = json.Unmarshal([]byte(resp.String()), &r)
	if err != nil {
		return SystemRouterboard{}, err
	}

	return r, nil
}

func (c *Client) GetIPFirewallAddresslist(query map[string]string) (IPFirewallAddresslists, error) {
	url := c.generateGetUrl("/ip/firewall/address-list", query)

	resp, err := c.GET(url)
	if err != nil {
		return IPFirewallAddresslists{}, err
	}

	r := IPFirewallAddresslists{}
	err = json.Unmarshal([]byte(resp.String()), &r)
	if err != nil {
		return IPFirewallAddresslists{}, err
	}

	return r, nil
}

func (c *Client) generateGetUrl(path string, query map[string]string) string {
	u, _ := url.Parse(fmt.Sprintf("%s%s", c.url, path))
	q, _ := url.ParseQuery(u.RawQuery)
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Client) PUT(url string, body interface{}) (gjson.Result, error) {
	resp, err := c.client.R().SetBody(body).Put(url)

	if err != nil {
		return gjson.Result{}, err
	} else if resp.RawResponse.StatusCode/200 != 1 {
		return gjson.Result{}, fmt.Errorf(fmt.Sprintf("Response status code %d, response: %s", resp.RawResponse.StatusCode, resp.Body()))
	}

	respParse := gjson.Parse(string(resp.Body()))

	if respParse.Get("detail").String() != "" {
		return respParse, fmt.Errorf(respParse.Get("detail").String())
	}

	return respParse, nil
}

func (c *Client) PutIPFirewallAddresslist(body IPFirewallAddresslist) (IPFirewallAddresslist, error) {
	url := fmt.Sprintf("%s/ip/firewall/address-list", c.url)

	resp, err := c.PUT(url, body)
	if err != nil {
		return IPFirewallAddresslist{}, err
	}

	r := IPFirewallAddresslist{}
	err = json.Unmarshal([]byte(resp.String()), &r)
	if err != nil {
		return IPFirewallAddresslist{}, err
	}

	return r, nil
}
