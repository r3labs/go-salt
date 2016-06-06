/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"encoding/json"
	"fmt"
)

// Client ...
type Client struct {
	Connector *Connector
}

// NewClient ...
func NewClient(config Config) (*Client, error) {
	c := Client{}
	c.Connector = NewConnector(config)
	err := c.Connector.Authenticate()

	return &c, err
}

// Minions ...
func (c *Client) Minions() (map[string]Minion, error) {
	m := MinionsResponse{}

	resp, err := c.Connector.Get("/minions")
	if err != nil {
		return m.Minions[0], err
	}

	data, err := parseResponse(resp)
	if err != nil {
		return m.Minions[0], err
	}

	err = json.Unmarshal(*data, &m)

	return m.Minions[0], err
}

// Minion ...
func (c *Client) Minion(id string) (Minion, error) {
	var m Minion

	uri := fmt.Sprintf("/minions/%s", id)
	resp, err := c.Connector.Get(uri)
	if err != nil {
		return m, err
	}

	data, err := parseResponse(resp)
	fmt.Println(string(*data))
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(*data, &m)

	return m, err
}

// Jobs ...
func (c *Client) Jobs() ([]map[string]Job, error) {
	jr := JobsResponse{}

	resp, err := c.Connector.Get("/jobs")
	if err != nil {
		return jr.Jobs, err
	}

	data, err := parseResponse(resp)
	if err != nil {
		return jr.Jobs, err
	}

	err = json.Unmarshal(*data, &jr)

	return jr.Jobs, err
}

// Job ...
func (c *Client) Job(id string) (Job, error) {
	j := JobResponse{}

	uri := fmt.Sprintf("/jobs/%s", id)
	resp, err := c.Connector.Get(uri)
	if err != nil {
		return Job{}, err
	}

	data, err := parseResponse(resp)
	if err != nil {
		return Job{}, err
	}

	err = json.Unmarshal(*data, &j)

	return j.Job[0], err
}

// Execute ...
func (c *Client) Execute(function, command, target, targetType string) (string, error) {
	er := ExecutionResponse{}

	req := fmt.Sprintf(`{"fun": "%s", "arg": "%s", "tgt": "%s", "expr_form": "%s"}`, function, command, target, targetType)

	resp, err := c.Connector.Post("/minions", []byte(req))
	if err != nil {
		return "", err
	}

	data, err := parseResponse(resp)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(*data, &er)

	return er.Job[0].ID, err
}
