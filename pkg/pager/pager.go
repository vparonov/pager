package pager

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/vparonov/pager/pkg/entities"
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
	template, ok := p.repository.FindIssueType(typeName)

	if !ok {
		return "", fmt.Errorf("%s issue type not found", typeName)
	}
	id := uuid.NewString()

	body := replacePlaceholders(template, placeHolderValues)

	err := p.repository.InsertIssue(&entities.Issue{
		ID:   id,
		Body: body,
	})

	if err != nil {
		return "", err
	}

	return id, nil
}

func (p *pager) ClearIssue(id string, userID string) error {
	return nil
}

func replacePlaceholders(template string, placeHolderValues map[string]string) string {
	// naive implementation of template replacement
	s := template
	for k, v := range placeHolderValues {
		s = strings.ReplaceAll(s, k, v)
	}

	return s
}
