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
