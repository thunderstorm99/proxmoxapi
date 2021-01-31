package main

import (
	"net/http"
	"strconv"
)

// GetNodes returns an array of nodes (hosts) for a given proxmox connection
func (p *ProxmoxConnection) GetNodes() ([]Node, error) {
	// data struct for API call
	var d struct {
		Data []Node `json:"data"`
	}

	// call API
	if err := p.callAPI("/nodes", http.MethodGet, &d); err != nil {
		return nil, err
	}

	// transform into proper array
	var nodes []Node
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
