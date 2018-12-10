package firecracker

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/go-resty/resty"
)

type InstanceState string

const (
	Uninitialized InstanceState = "Uninitialized"
	Starting      InstanceState = "Starting"
	Running       InstanceState = "Running"
	Halting       InstanceState = "Halting"
	Halted        InstanceState = "Halted"
)

type instanceInfo struct {
	ID    string        `json:"id"`
	State InstanceState `json:"state,omitempty"`
}

type apiError struct {
	Message string `json:"fault_message"`
}

var (
	errInvalidServerError    = errors.New("invalid server error response")
	errInvalidServerResponse = errors.New("invalid server response")
)

// Firecracker is a HTTP API client for firecracker.
// Provides both unix file socket and host:port connectivity.
type Firecracker struct {
	id     string
	client *resty.Client
}

// NewSocket creates a firecracker client instance, uses a unix socket file for communication.
func NewSocket(path string) *Firecracker {
	cracker := &Firecracker{
		client: resty.NewWithClient(&http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", path)
				},
			},
		}),
	}

	cracker.client.SetHostURL("http://localhost")

	_, err := cracker.State()
	if err != nil {
		return nil
	}

	return cracker
}

// New creates a firecracker client instance, uses a host:port combination for communication.
func New(host string, port int) *Firecracker {
	cracker := &Firecracker{
		client: resty.New(),
	}

	cracker.client.SetHostURL(fmt.Sprintf("http://%v:%v", host, port))

	_, err := cracker.State()
	if err != nil {
		return nil
	}

	return cracker
}

func (cracker *Firecracker) responseError(resp *resty.Response, status int, strict bool) error {
	if resp.StatusCode() != status {
		if e, ok := resp.Error().(*apiError); ok {
			return errors.New(e.Message)
		}

		return errInvalidServerError
	}

	if strict {
		return errInvalidServerResponse
	}

	return nil
}

func (cracker *Firecracker) responseErrorStrict(resp *resty.Response, status int) error {
	return cracker.responseError(resp, status, true)
}

func (cracker *Firecracker) responseErrorLoose(resp *resty.Response, status int) error {
	return cracker.responseError(resp, status, false)
}

// State returns instance ID.
func (cracker *Firecracker) ID() (string, error) {
	_, err := cracker.State()
	if err != nil {
		return "", err
	}

	return cracker.id, nil
}

// State returns the state of the instance, one of: Uninitialized, Starting, Running, Halting, Halted.
func (cracker *Firecracker) State() (InstanceState, error) {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&instanceInfo{}).
		SetError(&apiError{}).
		Get("/")

	if err != nil {
		return Uninitialized, err
	}

	if err = cracker.responseErrorLoose(resp, http.StatusOK); err != nil {
		return Uninitialized, err
	}

	if info, ok := resp.Result().(*instanceInfo); ok {
		cracker.id = info.ID
		return info.State, nil
	}

	return Uninitialized, errInvalidServerResponse
}

func (cracker *Firecracker) action(typ string, payload string) error {
	resp, err := cracker.client.R().
		SetHeader("Accept", "application/json").
		SetError(&apiError{}).
		SetBody(map[string]interface{}{
			"action_type": typ,
			"payload":     payload,
		}).
		Put("/")

	if err != nil {
		return err
	}


	return cracker.responseErrorStrict(resp, http.StatusNoContent)
}

// Rescan available block devices.
func (cracker *Firecracker) Rescan(payload string) error {
	return cracker.action("BlockDeviceRescan", payload)
}

// Start starts the firecracker instance.
func (cracker *Firecracker) Start(payload string) error {
	return cracker.action("InstanceStart", payload)
}

// Halt halts the firecracker instance, duh.
func (cracker *Firecracker) Halt(payload string) error {
	return cracker.action("InstanceHalt", payload)
}
