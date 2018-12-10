package firecracker

import (
	"net/http"
)

type NetworkInterface struct {
	ID string `json:"iface_id"`

	GuestMACAddr string `json:"guest_mac,omitempty"`
	HostDevName  string `json:"host_dev_name,omitempty"`
	// Device will reply to HTTP GET metadata requests
	AllowGettingMetadata bool `json:"allow_mmds_requests,omitempty"`

	RxLimiter RateLimiter `json:"rx_rate_limiter,omitempty"`
	TxLimiter RateLimiter `json:"Tx_rate_limiter,omitempty"`
}

// CreateNetworkInterface creates a network interface.
func (cracker *Firecracker) CreateNetworkInterface(id string, netInterface NetworkInterface) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetPathParams(map[string]string{
			"iface_id": id,
		}).
		SetError(&apiError{}).
		SetBody(netInterface).
		Patch("/network-interface/{iface_id}")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}
