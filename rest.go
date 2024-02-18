package gravity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// All error constants
var (
	ErrJSONUnmarshal = errors.New("json unmarshal")
)

var (
	// Marshal defines function used to encode JSON payloads
	Marshal func(v interface{}) ([]byte, error) = json.Marshal
	// Unmarshal defines function used to decode JSON payloads
	Unmarshal func(src []byte, v interface{}) error = json.Unmarshal
)

func unmarshal(data []byte, v interface{}) error {
	err := Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}

	return nil
}

type APIErrorMessage struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
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
	err := Unmarshal(body, &msg)
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

func (g *Gravity) RequestWithQueryParam(method string, endpoint string, params map[string]string, options ...RequestOption) (response []byte, err error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return
	}

	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return g.request(method, u.String(), "", nil, options...)
}

func (g *Gravity) RequestWithJSON(method string, endpoint string, data interface{}, options ...RequestOption) (response []byte, err error) {
	var body []byte
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	return g.request(method, endpoint, "application/json; charset=utf-8", body, options...)
}

func (g *Gravity) RequestWithFormURLEncoded(method string, endpoint string, data url.Values, options ...RequestOption) (response []byte, err error) {
	bodyReader := strings.NewReader(data.Encode())
	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return
	}

	return g.request(method, endpoint, "application/x-www-form-urlencoded; charset=utf-8", body, options...)
}

func (g *Gravity) request(method string, endpoint string, contentType string, b []byte, options ...RequestOption) (response []byte, err error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
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
		err = Unmarshal(response, &msg)
		if err != nil {
			return
		}
		if msg.ErrNo != ErrNoSuccess {
			err = newRestError(req, resp, response)
		}
	default:
		err = newRestError(req, resp, response)
	}

	return
}
