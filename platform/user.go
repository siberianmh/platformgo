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
	"time"
)

type UserService interface {
	Get(context.Context) (*User, error)
}

// @typescript-ignore UserServiceImpl
type UserServiceImpl struct {
	client *Client
}

// User represents a user on Platform.
type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	SiteAdmin bool   `json:"site_admin"`
	AvatarURL string `json:"avatar_url,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Get returns the currently authenticated user.
func (s *UserServiceImpl) Get(ctx context.Context) (*User, error) {
	var resp *User
	err := s.client.newRequest(ctx, "GET", "/users/@me", nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
