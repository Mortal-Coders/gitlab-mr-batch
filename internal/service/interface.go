package service

import "github.com/mortal-coders/gitlab-mr-batch/internal/model"

type GitlabClient interface {
	ListNamespaces() ([]model.GitlabNamespace, error)
	ListUsersInGroup(groupId int) ([]model.GitlabUser, error)
	ListProjectsInGroup(groupId int) ([]model.GitlabProject, error)
	CreateMergeRequest(req model.GitlabMergeRequest) (*model.GitlabMergeRequest, error)
}
