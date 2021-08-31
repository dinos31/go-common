package httphub

import (
	"io"
	"net/http"
)

// HTTPClients provides an interface to manage http.
type HTTPClients interface {
	Do(req *http.Request) (response *http.Response, err error)
}

// HTTPVars struct to set http client vars
type HTTPVars struct {
	Path         string
	Method       string
	Body         io.Reader
	Client       HTTPClients
	HeaderAPIKey string
}

// Do used to create http client
func (httpVars HTTPVars) Do(req *http.Request) (response *http.Response, err error) {
	response, err = httpVars.Client.Do(req)
	return
}

// Client creates Http client
var Client = HTTPVars{
	Client: &http.Client{},
}
