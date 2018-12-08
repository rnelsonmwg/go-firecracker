package firecracker

import (
	"errors"
	"net/http"
)

type bootSource struct {
	KernelImagePath string `json:"kernel_image_path"`
	BootArgs        string `json:"boot_args,omitempty"`
}

// SetBootSource creates or updates the boot source.
// Right now firecracker supports only LocalImage sources.
func (cracker *Firecracker) SetBootSource(imagePath string, bootArgs string) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(&bootSource{
			KernelImagePath: imagePath,
			BootArgs:        bootArgs,
		}).
		Put("/boot-source")

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
