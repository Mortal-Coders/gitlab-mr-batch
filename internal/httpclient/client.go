package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
	tokenConfig TokenConfig
	maxRetries  int
}

type Config struct {
	Timeout         time.Duration
	MaxIdleConns    int
	IdleConnTimeout time.Duration
	MaxRetries      int
	TokenConfig     TokenConfig
}

func NewClient(cfg Config) *Client {
	transport := &http.Transport{
		MaxIdleConns:        cfg.MaxIdleConns,
		IdleConnTimeout:     cfg.IdleConnTimeout,
		MaxIdleConnsPerHost: cfg.MaxIdleConns,
	}

	client := &http.Client{
		Timeout:   cfg.Timeout,
		Transport: transport,
	}

	return &Client{
		Client:      client,
		tokenConfig: cfg.TokenConfig,
		maxRetries:  cfg.MaxRetries,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	switch c.tokenConfig.Type {
	case Bearer:
		req.Header.Set("Authorization", "Bearer "+c.tokenConfig.Value)
	case Basic:
		req.Header.Set("Authorization", "Basic "+c.tokenConfig.Value)
	case APIKey:
		req.Header.Set("X-Api-Key", c.tokenConfig.Value)
	case Custom:
		if c.tokenConfig.HeaderName == "" {
			return nil, fmt.Errorf("custom header name is required for CustomHeader type")
		}
		req.Header.Set(c.tokenConfig.HeaderName, c.tokenConfig.Value)
	default:
		return nil, fmt.Errorf("unsupported token type: %s", c.tokenConfig.Type)
	}

	var resp *http.Response
	var err error

	// Retry mekanizması
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		resp, err = c.Client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		if attempt < c.maxRetries {
			time.Sleep(time.Millisecond * 100 * time.Duration(attempt+1))
		}
	}

	return resp, err
}
