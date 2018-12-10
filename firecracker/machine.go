package firecracker

import (
	"errors"
	"net/http"
)

type CPUTemplate string

const (
	C3Template CPUTemplate = "C3"
	T2Template CPUTemplate = "T2"
)

type MachineConfig struct {
	VCpuCount   int64       `json:"vcpu_count"`
	MemSizeMb   int64       `json:"mem_size_mib"`
	HTEnabled   bool        `json:"ht_enabled"`
	CPUTemplate CPUTemplate `json:"cpu_template"`
}

// Config returns machine config of the vm, cpu/mem limits
func (cracker *Firecracker) Config() (*MachineConfig, error) {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetResult(&MachineConfig{}).
		Get("/machine-config")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusNoContent {
		if e, ok := resp.Error().(*apiError); ok {
			return nil, errors.New(e.Message)
		}

		return nil, errInvalidServerError
	}

	if config, ok := resp.Result().(*MachineConfig); ok {
		return config, nil
	}

	return nil, errInvalidServerResponse
}

// SetConfig sets desired machine config for the vm, cpu/mem limits
func (cracker *Firecracker) SetConfig(config *MachineConfig) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(config).
		Put("/machine-config")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}
