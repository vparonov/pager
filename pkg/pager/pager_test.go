package pager

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vparonov/pager/pkg/repository"
)

func TestPager(t *testing.T) {
	repo := newMemoryRepo()

	p := New(repo)

	err := p.RegisterIssueType("test", "test @param1")

	assert.Nil(t, err)

	id, err := p.CreateIssue("test", map[string]string{"@param1": "paramValue1"})

	assert.Nil(t, err)

	assert.NotEqual(t, 0, len(id))
}

type inMemoryRepo struct {
	issueTypes map[string]string
	issues     map[string]string
}

func newMemoryRepo() repository.Repository {
	return &inMemoryRepo{
		issueTypes: make(map[string]string),
		issues:     make(map[string]string),
	}
}

func (r *inMemoryRepo) UpsertIssueType(typeName string, template string) error {
	r.issueTypes[typeName] = template
	return nil
}

func (r *inMemoryRepo) FindIssueType(typeName string) (string, bool) {
	template, ok := r.issueTypes[typeName]

	return template, ok
}

func (r *inMemoryRepo) InsertIssue(id string, body string) error {
	_, ok := r.issues[id]

	if ok {
		return fmt.Errorf("duplicated issue id = %s", id)
	}

	r.issues[id] = body
	return nil
}
