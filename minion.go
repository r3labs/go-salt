/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

// MinionsResponse ...
type MinionsResponse struct {
	Minions []map[string]Minion `json:"return"`
}

// Minion ...
type Minion struct {
	ID            string   `json:"id"`
	Name          string   `json:"nodename"`
	Host          string   `json:"host"`
	Domain        string   `json:"domain"`
	OS            string   `json:"os"`
	OSRelease     string   `json:"osrelease"`
	OSName        string   `json:"osfullname"`
	Kernel        string   `json:"kernel"`
	KernelRelease string   `json:"kernelrelease"`
	Shell         string   `json:"shell"`
	ARCH          string   `json:"osarch"`
	CPUS          int      `json:"num_cpus"`
	RAM           int      `json:"mem_total"`
	CPUModel      string   `json:"cpu_model"`
	CPUFlags      []string `json:"cpu_flags"`
	Virtual       string   `json:"virtual"`
	IPv4          []string `json:"ipv4"`
	IPv6          []string `json:"ipv6"`
	Path          string   `json:"path"`
	ServerID      int      `json:"server_id"`
}
