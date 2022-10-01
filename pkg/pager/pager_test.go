package pager

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vparonov/pager/pkg/entities"
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

	issue, found := repo.FindIssue(id)

	assert.Equal(t, true, found)
	assert.NotNil(t, issue)

	err = p.ClearIssue(id, "vangel@elastecad.com")

	assert.Nil(t, err)

	err = p.ClearIssue(id, "example@elastecad.com")

	assert.Nil(t, err)

	resolutions, err := repo.GetResolutions(id)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(resolutions))

}

type inMemoryRepo struct {
	issueTypes  map[string]string
	issues      map[string]string
	resolutions map[string]string
}

func newMemoryRepo() repository.Repository {
	return &inMemoryRepo{
		issueTypes:  make(map[string]string),
		issues:      make(map[string]string),
		resolutions: make(map[string]string),
	}
}

func (r *inMemoryRepo) Open() error {
	return nil
}

func (r *inMemoryRepo) Close() error {
	return nil
}

func (r *inMemoryRepo) UpsertIssueType(typeName string, template string) error {
	r.issueTypes[typeName] = template
	return nil
}

func (r *inMemoryRepo) FindIssueType(typeName string) (string, bool) {
	template, ok := r.issueTypes[typeName]

	return template, ok
}

func (r *inMemoryRepo) InsertIssue(issue *entities.Issue) error {
	_, ok := r.issues[issue.ID]

	if ok {
		return fmt.Errorf("duplicated issue id = %s", issue.ID)
	}

	issueBytes, _ := entities.IssueToJson(issue)
	r.issues[issue.ID] = string(issueBytes)
	return nil
}

func (r *inMemoryRepo) FindIssue(id string) (*entities.Issue, bool) {
	issueString, ok := r.issues[id]

	if !ok {
		return nil, false
	}

	issue, err := entities.NewIssueFromJson([]byte(issueString))

	if err != nil {
		return nil, false
	}

	return issue, true
}

func (r *inMemoryRepo) InsertResolution(resolution *entities.Resolution) error {
	resolutions, err := r.GetResolutions(resolution.IssueID)

	if err != nil {
		return err
	}

	resolutions = append(resolutions, resolution)

	jsonBytes, err := entities.ResolutionSliceToJson(resolutions)

	if err != nil {
		return err
	}

	r.resolutions[resolution.IssueID] = string(jsonBytes)

	return nil
}

func (r *inMemoryRepo) GetResolutions(issueID string) ([]*entities.Resolution, error) {
	resolutions, ok := r.resolutions[issueID]

	if !ok {
		return nil, nil
	}
	return entities.NewResolutionsSliceFromJson([]byte(resolutions))
}
