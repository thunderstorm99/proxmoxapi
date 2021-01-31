package proxmoxapi

import (
	"net/http"
)

// GetVMs gets all virtual machines, as well as containers for a given node
func (p *ProxmoxConnection) GetVMs(node string) ([]VM, error) {
	// data struct for API call
	var d struct {
		Data []VM `json:"data"`
	}

	// call API
	if err := p.callAPI("/nodes/"+node+"/qemu", http.MethodGet, &d); err != nil {
		return nil, err
	}

	// transform into proper array
	var vms []VM
	vms = append(vms, d.Data...)

	return vms, nil
}

// GetVersion returns version info for a given node
func (p *ProxmoxConnection) GetVersion(node string) (Version, error) {
	var d struct {
		Data Version `json:"data"`
	}

	if err := p.callAPI("/nodes/"+node+"/version", http.MethodGet, &d); err != nil {
		return d.Data, err
	}

	return d.Data, nil
}

// StartAll starts all containers and virtual machines on a given host
func (p *ProxmoxConnection) StartAll(node string) error {
	var d struct {
		Data string `json:"data"`
	}

	if err := p.callAPI("/nodes/"+node+"/startall?force=1", http.MethodPost, &d); err != nil {
		return err
	}

	return nil
}
