package firecracker

import (
	"net/http"
)

type CPUTemplate string

const (
	C3Template CPUTemplate = "C3"
	T2Template CPUTemplate = "T2"
)

type MachineConfig struct {
	VCpuCount   int64       `json:"vcpu_count,omitempty"`
	MemSizeMb   int64       `json:"mem_size_mib,omitempty"`
	HTEnabled   bool        `json:"ht_enabled,omitempty"`
	CPUTemplate CPUTemplate `json:"cpu_template,omitempty"`
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

	if err = cracker.responseErrorLoose(resp, http.StatusNoContent); err != nil {
		return nil, err
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
