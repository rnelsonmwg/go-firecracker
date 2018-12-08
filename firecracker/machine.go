package firecracker

import (
	"errors"
	"net/http"
)

type MachineConfig struct {
}

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

func (cracker *Firecracker) SetConfig(config *MachineConfig) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(config).
		Put("/machine-config")

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusNoContent {
		if e, ok := resp.Error().(*apiError); ok {
			return errors.New(e.Message)
		}

		return errInvalidServerError
	}

	return errInvalidServerResponse
}
