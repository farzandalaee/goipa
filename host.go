package ipa

import (
	"encoding/json"
)

// HostRecord encapsulates hbac data returned from ipa host commands
type HostRecord struct {
	Dn   string   `json:"dn"`
	FQDN []string `json:"fqdn"`
}

// HbacFind Fetch host list by call the FreeIPA host-find method
func (c *Client) HbacFind() ([]HostRecord, error) {
	options := map[string]interface{}{
		"no_members": false,
		"all":        true}

	res, err := c.rpc("host_find", []string{}, options)

	if err != nil {
		return nil, err
	}

	var hostRec []HostRecord
	err = json.Unmarshal(res.Result.Data, &hostRec)
	if err != nil {
		return nil, err
	}

	return hostRec, nil
}
