package routeros

import (
	"encoding/json"
	"fmt"
	"net/url"

	types "github.com/littleya/routeros-rest/types"
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

func (c *Client) GetStruct(s types.Types, query map[string]string) error {
	url := c.generateGetUrl(s.GetRestPath(), query)
	resp, err := c.client.R().Get(url)

	if err != nil {
		return err
	} else if resp.RawResponse.StatusCode/200 != 1 {
		return fmt.Errorf(fmt.Sprintf("Response status code %d, response: %s", resp.RawResponse.StatusCode, resp.Body()))
	}

	respParse := gjson.Parse(string(resp.Body()))

	if respParse.Get("detail").String() != "" {
		return fmt.Errorf(respParse.Get("detail").String())
	}

	err = json.Unmarshal([]byte(respParse.String()), s)
	if err != nil {
		return err
	}

	return nil
}

// func (c *Client) GetSystemRouterboard() (types.SystemRouterboard, error) {
// 	url := c.generateGetUrl("/system/routerboard", map[string]string{})
// 	r := types.SystemRouterboard{}
// 	if err := c.GetStruct(&r, query); err != nil {
// 		return r, err
// 	}
// 	return r, nil
// }

func (c *Client) GetIPFirewallAddresslist(query map[string]string) (types.IPFirewallAddresslists, error) {
	url := c.generateGetUrl("/ip/firewall/address-list", query)

	resp, err := c.GET(url)
	if err != nil {
		return types.IPFirewallAddresslists{}, err
	}

	r := types.IPFirewallAddresslists{}
	err = json.Unmarshal([]byte(resp.String()), &r)
	if err != nil {
		return types.IPFirewallAddresslists{}, err
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

func (c *Client) PutStruct(s types.Types, body interface{}) error {
	url := fmt.Sprintf("%s%s", c.url, s.GetRestPath())

	resp, err := c.client.R().SetBody(body).Put(url)

	if err != nil {
		return err
	} else if resp.RawResponse.StatusCode/200 != 1 {
		return fmt.Errorf(fmt.Sprintf("Response status code %d, response: %s", resp.RawResponse.StatusCode, resp.Body()))
	}

	respParse := gjson.Parse(string(resp.Body()))

	if respParse.Get("detail").String() != "" {
		return fmt.Errorf(respParse.Get("detail").String())
	}

	err = json.Unmarshal([]byte(resp.String()), s)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PutIPFirewallAddresslist(body types.IPFirewallAddresslist) (types.IPFirewallAddresslist, error) {
	url := fmt.Sprintf("%s/ip/firewall/address-list", c.url)

	resp, err := c.PUT(url, body)
	if err != nil {
		return types.IPFirewallAddresslist{}, err
	}

	r := types.IPFirewallAddresslist{}
	err = json.Unmarshal([]byte(resp.String()), &r)
	if err != nil {
		return types.IPFirewallAddresslist{}, err
	}

	return r, nil
}
