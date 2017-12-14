package api

import (
	"net/http"
	"time"

	"github.com/screwdriver-cd/sd-cmd/config"
)

const (
	timeoutSec = 10
)

// API is a Screwdriver API endpoint
type API interface {
	SetJWT() error
	GetCommand(namespace, command, version string) (*Command, error)
}

type client struct {
	baseURL  string
	apiToken string
	jwt      string
	client   *http.Client
}

// JWT is a Screwdriver JWT
type JWT struct {
	Token string `json:"token"`
}

// ResponseError is an error response from the Screwdriver API
type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Reason     string `json:"error"`
	Message    string `json:"message"`
}

// Command is a Screwdriver Command
type Command struct {
	Namespace   string `json:"namespace"`
	Command     string `json:"command"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Format      string `json:"format"`
	Habitat     struct {
		Mode    string `json:"mode"`
		Package string `json:"package"`
		Binary  string `json:"binary"`
	} `json:"habitat"`
	Docker struct {
		Image string `json:"image"`
	} `json:"docker"`
	Binary struct {
		File string `json:"file"`
	} `json:"binary"`
}

// New returns API object
func New() (API, error) {
	c := &client{
		baseURL:  config.SDAPIURL,
		apiToken: config.SDAPIToken,
		client:   &http.Client{Timeout: timeoutSec * time.Second},
	}
	return API(c), nil
}

// SetJWT sets jwt to API
func (c *client) SetJWT() error {
	return nil
}

// GetCommand returns Command from Screwdriver API
func (c client) GetCommand(namespace, command, version string) (*Command, error) {
	cmd := new(Command)
	cmd.Namespace = namespace
	cmd.Command = command
	cmd.Version = version
	cmd.Format = "binary"
	return cmd, nil
}