package http

import (
	"github.com/ess/hype"

	"github.com/ess/fbz/pkg/fbz"
)

// Driver is an object that knows specifically how to interact with the
// Engine Yard API at the HTTP level
type Driver struct {
	raw         *hype.Driver
	token       string
	accept      *hype.Header
	contentType *hype.Header
	userAgent   *hype.Header
}

// NewDriver takes a base URL for an Engine Yard API and a token, returning a
// Driver that can be used to interact with the API in question.
func NewDriver(baseURL string, token string) (*Driver, error) {
	raw, err := hype.New(baseURL)
	if err != nil {
		return nil, err
	}

	d := &Driver{
		raw,
		token,
		hype.NewHeader("Accept", "application/json"),
		hype.NewHeader("Content-Type", "application/json"),
		hype.NewHeader("User-Agent", "fbz/0.1.0 (https://github.com/ess/fbz)"),
	}

	return d, nil
}

func (driver *Driver) Token() string {
	return driver.token
}

// Post performs a POST operation for the given path, params, and data against
// the upstream API. it returns a byte array and an error.
func (driver *Driver) Post(path string, data []byte) fbz.Response {
	response := driver.
		raw.
		Post(path, nil, data).
		WithHeaderSet(driver.accept, driver.contentType, driver.userAgent).
		Response()

	return response
}

/*
Copyright 2021 Dennis Walters

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
