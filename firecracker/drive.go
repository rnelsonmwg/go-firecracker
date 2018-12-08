package firecracker

import (
	"errors"
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

func (cracker *Firecracker) SetDrive(id string, hostPath string, isReadOnly bool, isRoot bool, uuid string) error {
	return cracker.SetDriveWithLimiter(id, hostPath, isReadOnly, isRoot, uuid, nil)
}

func (cracker *Firecracker) SetDriveWithLimiter(id string, hostPath string, isReadOnly bool, isRoot bool, uuid string, limiter *RateLimiter) error {
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

	if resp.StatusCode() != http.StatusNoContent {
		if e, ok := resp.Error().(*apiError); ok {
			return errors.New(e.Message)
		}

		return errInvalidServerError
	}

	return errInvalidServerResponse
}

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

	if resp.StatusCode() != http.StatusNoContent {
		if e, ok := resp.Error().(*apiError); ok {
			return errors.New(e.Message)
		}

		return errInvalidServerError
	}

	return errInvalidServerResponse
}
