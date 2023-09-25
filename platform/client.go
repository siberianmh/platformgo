// Copyright 2023 Siberian Syndicate, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://api-platform.siberianmh.com"
	userAgent      = "platform-go/1.0"
)

// @typescript-ignore Client
type Client struct {
	apiKey     string
	httpClient *http.Client
	endpoint   string
	userAgent  string

	WebAnalytics WebAnalyticsService
	User         UserService
}

func NewClient(key string) *Client {
	c := &Client{
		apiKey: key,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		endpoint:  defaultBaseURL,
		userAgent: userAgent,
	}

	c.WebAnalytics = &WebAnalyticsServiceImpl{client: c}
	c.User = &UserServiceImpl{client: c}

	return c
}

func (c *Client) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *Client) SetToken(apiKey string) {
	c.apiKey = apiKey
}

func (c *Client) newRequest(ctx context.Context, method, path string, body, resp any) error {
	var bodyBuf io.ReadWriter
	if body != nil {
		bodyBuf = &bytes.Buffer{}
		enc := json.NewEncoder(bodyBuf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.endpoint, path), bodyBuf)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	if c.userAgent != "" {
		req.Header.Add("User-Agent", c.userAgent)
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("response not successful status=%d", res.StatusCode)
	}

	if err := json.NewDecoder(&buf).Decode(&resp); err != nil {
		return fmt.Errorf("%s: %w", "decoding response", err)
	}

	return nil
}
