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
	"context"
	"fmt"
	"net/http"
	"time"
)

type OrgsService interface {
	List(ctx context.Context) ([]*Organization, error)
	Get(ctx context.Context, org string) (*Organization, error)
}

// @typescript-ignore OrgsServiceImpl
type OrgsServiceImpl struct {
	client *Client
}

type Organization struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *OrgsServiceImpl) List(ctx context.Context) ([]*Organization, error) {
	var resp []*Organization

	err := s.client.newRequest(ctx, http.MethodGet, "/users/@me/orgs", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *OrgsServiceImpl) Get(ctx context.Context, org string) (*Organization, error) {
	var resp *Organization

	url := fmt.Sprintf("/orgs/%s", org)
	err := s.client.newRequest(ctx, http.MethodGet, url, nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
