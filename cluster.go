package main

import (
	"net/http"
	"strconv"
)

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

func (p *ProxmoxConnection) nextID() (int, error) {
	var d struct {
		Data string `json:"data"`
	}

	// call API
	if err := p.callAPI("/cluster/nextid", http.MethodGet, &d); err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(d.Data)
	if err != nil {
		return 0, err
	}

	return id, nil
}
