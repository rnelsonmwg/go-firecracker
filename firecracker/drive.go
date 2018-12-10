package firecracker

import (
	"net/http"
)

type StorageDrive struct {
	ID            string      `json:"drive_id"`
	HostPath      string      `json:"path_on_host"`
	IsRoot        bool        `json:"is_root_device"`
	IsReadOnly    bool        `json:"is_read_only"`
	PartitionUUID string      `json:"partuuid"`
	Limiter       RateLimiter `json:"rate_limiter"`
}

// SetDrive adds or updates a vm drive with specified ID.
func (cracker *Firecracker) SetDrive(id string, hostPath string, isReadOnly bool, isRoot bool, uuid string) error {
	return cracker.SetDriveWithLimiter(id, hostPath, isReadOnly, isRoot, uuid, nil)
}

func (cracker *Firecracker) SetDriveWithLimiter(
	id string,
	hostPath string,
	isReadOnly bool,
	isRoot bool,
	uuid string,
	limiter *RateLimiter,
) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetPathParams(map[string]string{
			"drive_id": id,
		}).
		SetError(&apiError{}).
		SetBody(&StorageDrive{
			ID:            id,
			HostPath:      hostPath,
			IsRoot:        isRoot,
			IsReadOnly:    isReadOnly,
			PartitionUUID: uuid,
			Limiter:       *limiter,
		}).
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
