package firecracker

import (
	"net/http"
)

// CreateMetadata adds custom vm related metadata.
func (cracker *Firecracker) CreateMetadata(metadata map[string]interface{}) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(metadata).
		Put("/mmds")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}

// UpdateMetadata updates custom vm related metadata.
func (cracker *Firecracker) UpdateMetadata(metadata map[string]interface{}) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(metadata).
		Patch("/mmds")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}

// Metadata returns custom vm related metadata.
func (cracker *Firecracker) Metadata() (map[string]interface{}, error) {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetResult(make(map[string]interface{})).
		Get("/mmds")

	if err != nil {
		return nil, err
	}

	if err = cracker.responseError(resp, http.StatusNoContent); err != nil {
		return nil, err
	}

	if metadata, ok := resp.Result().(map[string]interface{}); ok {
		return metadata, nil
	}

	return nil, errInvalidServerResponse
}
