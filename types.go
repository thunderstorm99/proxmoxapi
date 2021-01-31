package main

type node struct {
	ID             string `json:"id"`
	Level          string `json:"level"`
	Node           string `json:"node"`
	SSLFingerprint string `json:"ssl_fingerprint"`
	Status         string `json:"status"`
	Type           string `json:"type"`
}

type version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}

type vm struct {
	CPULoad   float64 `json:"cpu"`
	CPUs      int     `json:"cpus"`
	Disk      int     `json:"disk"`
	DiskRead  int     `json:"diskread"`
	DiskWrite int     `json:"diskwrite"`
	MaxDisk   int     `json:"maxdisk"`
	MaxMem    int     `json:"maxmem"`
	Mem       int     `json:"mem"`
	Name      string  `json:"name"`
	NetIn     int     `json:"netin"`
	NetOut    int     `json:"netout"`
	Node      string  `json:"node"`
	PID       string  `json:"pid"`
	Status    string  `json:"status"`
	Template  string  `json:"template"`
	Uptime    int     `json:"uptime"`
	VMID      string  `json:"vmid"`
}

// ProxmoxConnection holds the URL and authentication token for Proxmox
type ProxmoxConnection struct {
	Insecure bool
	Token    string
	URL      string
}
