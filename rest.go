package gravity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIErrorMessage struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// RESTError stores error information about a request with a bad response code.
// Message is not always present, there are cases where api calls can fail
// without returning a json message.
type RESTError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte

	Message *APIErrorMessage // Message may be nil.
}

func (e *RESTError) Error() string {
	return fmt.Sprintf("request failed: %s %s, response: %s", e.Request.Method, e.Request.URL, e.ResponseBody)
}

// newRestError returns a new REST API error.
func newRestError(req *http.Request, resp *http.Response, body []byte) *RESTError {
	restErr := &RESTError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
	}

	// Attempt to decode the error and assume no message was provided if it fails
	var msg *APIErrorMessage
	err := json.Unmarshal(body, &msg)
	if err == nil {
		restErr.Message = msg
	}

	return restErr
}

// RequestConfig is an HTTP request configuration.
type RequestConfig struct {
	Request                *http.Request
	ShouldRetryOnRateLimit bool
	MaxRestRetries         int
	Client                 *http.Client
}

// newRequestConfig returns a new HTTP request configuration based on parameters in Gravity.
func newRequestConfig(g *Gravity, req *http.Request) *RequestConfig {
	return &RequestConfig{
		ShouldRetryOnRateLimit: g.ShouldRetryOnRateLimit,
		MaxRestRetries:         g.MaxRestRetries,
		Client:                 g.Client,
		Request:                req,
	}
}

// RequestOption is a function which mutates request configuration.
// It can be supplied as an argument to any REST method.
type RequestOption func(cfg *RequestConfig)

func (g *Gravity) Request(method string, url string, data interface{}, options ...RequestOption) (response []byte, err error) {
	var body []byte
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	return g.request(method, url, body, options...)
}

func (g *Gravity) request(method string, url string, b []byte, options ...RequestOption) (response []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	cfg := newRequestConfig(g, req)
	for _, opt := range options {
		opt(cfg)
	}
	req = cfg.Request

	resp, err := g.Client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			fmt.Println("error closing resp body")
		}
	}()

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case http.StatusOK:
		// Check response errno and if not 0, make new rest error
		var msg *APIErrorMessage
		err = json.Unmarshal(response, &msg)
		if err != nil {
			return
		}
		if msg.Errno != 0 {
			err = newRestError(req, resp, response)
		}
	default:
		err = newRestError(req, resp, response)
	}

	return
}
