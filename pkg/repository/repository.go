package repository

import "github.com/vparonov/pager/pkg/entities"

// NB! All functions should be thread safe
type Repository interface {
	Open() error
	Close() error
	UpsertIssueType(typeName string, template string) error
	FindIssueType(typeName string) (string, bool)

	InsertIssue(issue *entities.Issue) error
	FindIssue(id string) (*entities.Issue, bool)

	InsertResolution(resolution *entities.Resolution) error
	GetResolutions(issueID string) ([]*entities.Resolution, error)
}
