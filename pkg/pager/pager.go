package pager

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vparonov/pager/pkg/repository"
)

type Pager interface {
	RegisterIssueType(typeName string, issueTemplate string) error
	CreateIssue(typeName string, placeHolderValues map[string]string) (string, error)
	ClearIssue(id string, userID string) error
}

type pager struct {
	repository repository.Repository
}

func New(repository repository.Repository) Pager {
	return &pager{repository: repository}
}

func (p *pager) RegisterIssueType(typeName string, issueTemplate string) error {
	return p.repository.UpsertIssueType(typeName, issueTemplate)
}

func (p *pager) CreateIssue(typeName string, placeHolderValues map[string]string) (string, error) {
	_, ok := p.repository.FindIssueType(typeName)

	if !ok {
		return "", fmt.Errorf("%s issue type not found", typeName)
	}
	id := uuid.NewString()
	return id, nil
}

func (p *pager) ClearIssue(id string, userID string) error {
	return nil
}
