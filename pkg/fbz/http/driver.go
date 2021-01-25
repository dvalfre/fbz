package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ess/fbz/pkg/fbz"
)

// Driver is an object that knows specifically how to interact with the
// Engine Yard API at the HTTP level
type Driver struct {
	raw     *http.Client
	baseURL url.URL
	token   string
}

// NewDriver takes a base URL for an Engine Yard API and a token, returning a
// Driver that can be used to interact with the API in question.
func NewDriver(baseURL string, token string) (*Driver, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	d := &Driver{
		&http.Client{Timeout: 20 * time.Second},
		*url,
		token,
	}

	return d, nil
}

func (driver *Driver) Token() string {
	return driver.token
}

// Post performs a POST operation for the given path, params, and data against
// the upstream API. it returns a byte array and an error.
func (driver *Driver) Post(path string, params fbz.Params, data []byte) fbz.Response {
	return driver.makeRequest("POST", path, paramsToValues(params), data)
}

func (driver *Driver) rawRequest(verb string, path string, params url.Values, data []byte) (*http.Response, []byte, error) {

	request, err := driver.newRequest(verb, path, params, data)
	if err != nil {
		return nil, nil, err
	}

	response, err := driver.raw.Do(request)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		fmt.Println("DEBUG - Got the following HTTP status:", response.StatusCode)
		return nil, nil,
			fmt.Errorf(
				"The upstream API returned the following status: %d",
				response.StatusCode,
			)
	}

	return response, body, nil
}

func (driver *Driver) makeRequest(verb string, path string, params url.Values, data []byte) fbz.Response {

	_, body, err := driver.rawRequest(verb, path, params, data)
	if err != nil {
		return fbz.Response{nil, err}
	}

	return fbz.Response{body, nil}
}

func (driver *Driver) newRequest(verb string, path string, params url.Values, data []byte) (*http.Request, error) {
	request, err := http.NewRequest(
		verb,
		driver.constructRequestURL(path, params),
		bytes.NewReader(data),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "fbz/0.1.0 (https://github.com/ess/fbz)")

	return request, nil
}

func (driver *Driver) constructRequestURL(path string, params url.Values) string {

	pathParts := []string{driver.baseURL.Path, path}

	requestURL := url.URL{
		Scheme:   driver.baseURL.Scheme,
		Host:     driver.baseURL.Host,
		Path:     strings.Join(pathParts, "/"),
		RawQuery: params.Encode(),
	}

	result := requestURL.String()

	return result
}

/*
Copyright 2018 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
