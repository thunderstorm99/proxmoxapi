package main

import "net/http"

func (p *ProxmoxConnection) getNodes() ([]node, error) {
	// data struct for API call
	var d struct {
		Data []node `json:"data"`
	}

	// call API
	if err := p.callAPI("/nodes", http.MethodGet, &d); err != nil {
		return nil, err
	}

	// transform into proper array
	var nodes []node
	nodes = append(nodes, d.Data...)

	return nodes, nil
}
