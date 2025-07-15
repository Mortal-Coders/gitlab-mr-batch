package model

import "time"

type GitlabUser struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	State       string    `json:"state"`
	Locked      bool      `json:"locked"`
	AvatarUrl   string    `json:"avatar_url"`
	WebUrl      string    `json:"web_url"`
	AccessLevel int       `json:"access_level"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
