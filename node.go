package main

import "net/http"

func (p *ProxmoxConnection) getVMs(node string) ([]vm, error) {
	// data struct for API call
	var d struct {
		Data []vm `json:"data"`
	}

	// call API
	if err := p.callAPI("/nodes/"+node+"/qemu", http.MethodGet, &d); err != nil {
		return nil, err
	}

	// transform into proper array
	var vms []vm
	vms = append(vms, d.Data...)

	return vms, nil
}
