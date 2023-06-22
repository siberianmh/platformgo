package platform

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type WebAnalyticsService interface {
	List(ctx context.Context, org string) ([]AnalyticsWebsite, error)
	Get(ctx context.Context, org string, domain string) (*AnalyticsWebsite, error)
	Create(context.Context, string, *CreateAnalyticsWebsiteRequest) (*AnalyticsWebsite, error)
}

// @typescript-ignore WebAnalyticsServiceImpl
type WebAnalyticsServiceImpl struct {
	client *Client
}

// AnalyticsWebsite represents an Web Analytics on Platform
type AnalyticsWebsite struct {
	ID        int64     `json:"id"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAnalyticsWebsiteRequest struct {
	Domain string `json:"domain"`
}

func (s *WebAnalyticsServiceImpl) List(ctx context.Context, org string) ([]AnalyticsWebsite, error) {
	var resp []AnalyticsWebsite
	url := fmt.Sprintf("/orgs/%s/analytics", org)
	err := s.client.newRequest(ctx, http.MethodGet, url, nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *WebAnalyticsServiceImpl) Get(ctx context.Context, org string, domain string) (*AnalyticsWebsite, error) {
	var resp *AnalyticsWebsite
	url := fmt.Sprintf("/orgs/%s/analytics/%s", org, domain)
	err := s.client.newRequest(ctx, http.MethodGet, url, nil, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *WebAnalyticsServiceImpl) Create(ctx context.Context, org string, createRequest *CreateAnalyticsWebsiteRequest) (*AnalyticsWebsite, error) {
	var resp *AnalyticsWebsite
	url := fmt.Sprintf("/orgs/%s/analytics", org)
	err := s.client.newRequest(ctx, http.MethodPost, url, createRequest, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
