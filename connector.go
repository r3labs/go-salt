/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config ...
type Config struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Debug         bool
	SSLSkipVerify bool
}

// Connector ...
type Connector struct {
	Config    Config
	Client    *http.Client
	AuthToken string
}

// NewConnector ...
func NewConnector(config Config) *Connector {
	c := Connector{Config: config}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.SSLSkipVerify,
		},
	}
	c.Client = &http.Client{Transport: tr}
	return &c
}

// Authenticate ...
func (c *Connector) Authenticate() error {
	url := fmt.Sprintf("https://%s:%s/login", c.Config.Host, c.Config.Port)

	data := fmt.Sprintf(`{ "username":"%s", "password":"%s", "eauth": "pam" }`, c.Config.Username, c.Config.Password)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to authenticate")
	}

	c.AuthToken = resp.Header.Get("X-Auth-Token")
	return nil
}

// Get ...
func (c *Connector) Get(uri string) (*http.Response, error) {
	url := fmt.Sprintf("https://%s:%s%s", c.Config.Host, c.Config.Port, uri)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.AuthToken)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to make request")
	}

	return resp, nil
}

// Post ...
func (c *Connector) Post(uri string, data []byte) (*http.Response, error) {
	url := fmt.Sprintf("https://%s:%s%s", c.Config.Host, c.Config.Port, uri)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.AuthToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 202 {
		return nil, errors.New("failed to make request")
	}

	return resp, nil
}

func parseResponse(resp *http.Response) (*[]byte, error) {
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return &data, err
}
