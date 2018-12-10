package firecracker

import (
	"net/http"
)

type StorageDrive struct {
	ID            string      `json:"drive_id"`
	HostPath      string      `json:"path_on_host"`
	IsRoot        bool        `json:"is_root_device"`
	IsReadOnly    bool        `json:"is_read_only"`
	PartitionUUID string      `json:"partuuid,omitempty"`
	Limiter       *RateLimiter `json:"rate_limiter,omitempty"`
}

func (cracker *Firecracker) SetDrive(
	drive *StorageDrive,
) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetPathParams(map[string]string{
			"drive_id": drive.ID,
		}).
		SetError(&apiError{}).
		SetBody(drive).
		Put("/drives/{drive_id}")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}

// UpdateDrivePath updates vm drive path by ID.
func (cracker *Firecracker) UpdateDrivePath(id string, hostPath string) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetPathParams(map[string]string{
			"drive_id": id,
		}).
		SetError(&apiError{}).
		SetBody(map[string]interface{}{
			"id":           id,
			"path_on_host": hostPath,
		}).
		Patch("/drives/{drive_id}")

	if err != nil {
		return err
	}

	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}
