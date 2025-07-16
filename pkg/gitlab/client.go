package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mortal-coders/gitlab-mr-batch/internal/httpclient"
	"github.com/mortal-coders/gitlab-mr-batch/internal/model"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client  *httpclient.Client
	baseUrl string
}

func NewGitlabClient(url, token string) *Client {
	cfg := httpclient.Config{
		Timeout:         10 * time.Second,
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
		MaxRetries:      3,
		TokenConfig: httpclient.TokenConfig{
			Type:       httpclient.Custom,
			Value:      token,
			HeaderName: "PRIVATE-TOKEN",
		},
	}

	client := httpclient.NewClient(cfg)

	return &Client{
		client:  client,
		baseUrl: url,
	}
}

func (c *Client) ListNamespaces() ([]model.GitlabNamespace, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v4/namespaces", c.baseUrl), nil)
	//request, err := http.NewRequest("GET", fmt.Sprintf("%s/namespaces.json", c.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	var bodyBuf bytes.Buffer
	reader := io.TeeReader(resp.Body, &bodyBuf)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d body: %s", resp.StatusCode, bodyBuf.String())
	}

	var list []model.GitlabNamespace
	if err = json.NewDecoder(reader).Decode(&list); err != nil {
		return nil, fmt.Errorf("unmarshal error (status: %d): %w | body: %q", resp.StatusCode, err, bodyBuf.String())
	}

	return list, nil
}

func (c *Client) ListUsersInGroup(groupId int) ([]model.GitlabUser, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/groups/%d/members", c.baseUrl, groupId), nil)
	//request, err := http.NewRequest("GET", fmt.Sprintf("%s/users.json", c.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	var bodyBuf bytes.Buffer
	reader := io.TeeReader(resp.Body, &bodyBuf)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d body: %s", resp.StatusCode, bodyBuf.String())
	}

	var list []model.GitlabUser
	if err = json.NewDecoder(reader).Decode(&list); err != nil {
		return nil, fmt.Errorf("unmarshal error (status: %d): %w | body: %q", resp.StatusCode, err, bodyBuf.String())
	}
	return list, nil
}

func (c *Client) ListProjectsInGroup(groupId int) ([]model.GitlabProject, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/groups/%d/projects", c.baseUrl, groupId), nil)
	//request, err := http.NewRequest("GET", fmt.Sprintf("%s/projects.json", c.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	var bodyBuf bytes.Buffer
	reader := io.TeeReader(resp.Body, &bodyBuf)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d body: %s", resp.StatusCode, bodyBuf.String())
	}

	var list []model.GitlabProject
	if err = json.NewDecoder(reader).Decode(&list); err != nil {
		return nil, fmt.Errorf("unmarshal error (status: %d): %w | body: %q", resp.StatusCode, err, bodyBuf.String())
	}
	return list, nil
}

func (c *Client) CreateMergeRequest(req model.GitlabMergeRequest) (*model.GitlabMergeRequest, error) {
	//TODO implement me
	panic("implement me")
}
