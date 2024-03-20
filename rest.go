package gravity

import (
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

func unmarshal(data interface{}, v interface{}) error {
	byteArr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = Unmarshal(byteArr, v)
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

func (g *Gravity) requestWithQuery(method, path string, queryParams interface{}) (response *APIErrorMessage, err error) {
	return g.request(method, path, "", queryParams)
}

func (g *Gravity) requestWithForm(method, path string, body interface{}) (response *APIErrorMessage, err error) {
	return g.request(method, path, "application/x-www-form-urlencoded; charset=utf-8", body)
}

func (g *Gravity) request(method, endpoint, contentType string, requestData interface{}) (response *APIErrorMessage, err error) {
	g.state.device.Timestamp = getstrts(time.Now().Unix())

	di := structToMapWithJSON(g.state.device)
	rd := structToMapWithJSON(requestData)

	// Merge di and rd
	for k, v := range di {
		rd[k] = v
	}

	if g.state.cred.Token != "" {
		rd["token"] = g.state.cred.Token
	}

	rd["sign"], err = generateSignature(rd)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	switch contentType {
	case "":
		params := url.Values{}
		for k, v := range rd {
			params.Add(k, v)
		}
		req, err = http.NewRequest(method, (endpoint + "?" + params.Encode()), nil)
	case "application/x-www-form-urlencoded; charset=utf-8":
		formData := url.Values{}
		for k, v := range rd {
			formData.Add(k, v)
		}
		req, err = http.NewRequest(method, endpoint, strings.NewReader(formData.Encode()))
	default:
		return nil, fmt.Errorf("invalid content type %s", contentType)
	}

	if err != nil {
		return
	}

	req.Header.Set("Host", "api.gravity.place")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("User-Agent", "okhttp/3.12.13")

	resp, err := g.client.Do(req)
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

	err = Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrNo != ErrNoSuccess {
		err = newRestError(req, resp, respBody)
		return response, err
	}

	return response, nil
}
