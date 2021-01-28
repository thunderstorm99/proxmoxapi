package main

import (
	"log"

	"github.com/thunderstorm99/apihandler"
)

// NewProxmoxConnection creates a new connection to Proxmox via url and token
func NewProxmoxConnection(url string, token string) ProxmoxConnection {
	return ProxmoxConnection{URL: url, Token: token}
}

func (p *ProxmoxConnection) callAPI(url string, method string, resp interface{}) error {
	api := apihandler.APICall{URL: p.URL + url, Method: method, Header: map[string]string{"Authorization": "PVEAPIToken=" + p.Token, "Accept": "application/json"}, Insecure: p.Insecure}

	log.Println("calling URL", api.URL)
	err := api.Exec(&resp)
	if err != nil {
		return err
	}

	return nil
}
