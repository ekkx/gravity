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
	"time"
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

func (g *Gravity) requestWithQuery(method, path string, queryParams, st interface{}) (response interface{}, err error) {
	return g.request(method, path, "", queryParams, st)
}

func (g *Gravity) requestWithJSON(method, path string, body, st interface{}) (response interface{}, err error) {
	return g.request(method, path, "application/json", body, st)
}

func (g *Gravity) requestWithForm(method, path string, body, st interface{}) (response interface{}, err error) {
	return g.request(method, path, "application/x-www-form-urlencoded", body, st)
}

func (g *Gravity) request(method, endpoint, contentType string, requestData, st interface{}) (r interface{}, err error) {
	g.state.device.IDFA = encrypt(g.state.cred.GAID)
	g.state.device.UWD = encrypt(g.state.cred.UUID)
	g.state.device.Timestamp = getstrts(time.Now().Unix())

	g.state.device.Sign, _ = generateSignature(structToMapWithJSON(g.state.device))

	di := structToMapWithJSON(g.state.device)
	rd := structToMapWithJSON(requestData)

	// Merge di and rd
	for k, v := range di {
		rd[k] = v
	}

	var req *http.Request

	switch contentType {
	case "":
		params := url.Values{}
		for k, v := range rd {
			params.Add(k, v)
		}

		req, err = http.NewRequest(method, (endpoint + "?" + params.Encode()), nil)

	case "application/json":
		jsonData, err := json.Marshal(rd)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json; charset=utf-8")

	case "application/x-www-form-urlencoded":
		formData := url.Values{}
		for k, v := range rd {
			formData.Add(k, v)
		}

		req, err = http.NewRequest(method, endpoint, strings.NewReader(formData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	default:
		return nil, fmt.Errorf("invalid content type %s", contentType)
	}

	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err2 := resp.Body.Close()
		if err2 != nil {
			fmt.Println("error closing resp body")
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var response *APIErrorMessage
	err = Unmarshal(respBody, &response)
	if err != nil {
		return
	}

	if response.ErrNo != ErrNoSuccess {
		err = newRestError(req, resp, respBody)
		return nil, err
	}

	return response.Data, nil
}
