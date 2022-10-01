package repository

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vparonov/pager/pkg/entities"
)

func TestBoltRepository(t *testing.T) {
	defer os.Remove("test.db")

	repo := NewBoltRepository("test.db")

	err := repo.Open()

	assert.Nil(t, err)

	err = repo.UpsertIssueType("test", "test")

	assert.Nil(t, err)

	template, found := repo.FindIssueType("test")

	assert.Equal(t, true, found)

	assert.Equal(t, "test", template)

	_, found = repo.FindIssueType("issue_type_that_does_not_exists")

	assert.Equal(t, false, found)

	err = repo.Close()

	assert.Nil(t, err)
}

func TestBoltRepositoryErrors(t *testing.T) {
	defer os.Remove("test.db")

	repo := NewBoltRepository("test.db")

	err := repo.Open()

	assert.Nil(t, err)

	// look for a type before it is created
	_, found := repo.FindIssueType("test")

	assert.Equal(t, false, found)
	err = repo.Close()

	assert.Nil(t, err)

}

func TestBoltRepositoryInsertIssue(t *testing.T) {
	defer os.Remove("test.db")

	repo := NewBoltRepository("test.db")

	err := repo.Open()

	assert.Nil(t, err)

	issue := &entities.Issue{
		ID:        "test",
		Body:      "test body", //
		CreatedAt: time.Now(),
	}

	err = repo.InsertIssue(issue) //

	assert.Nil(t, err)

	foundIssue, found := repo.FindIssue("test")

	assert.Equal(t, true, found)
	assert.Equal(t, "test body", foundIssue.Body)
	err = repo.Close()

	assert.Nil(t, err)

}

func TestMarshalArray(t *testing.T) {
	r1 := &entities.Resolution{
		IssueID: "1",
		UserID:  "vangel@elastecad.com",
		Ts:      time.Now(),
	}
	r2 := &entities.Resolution{
		IssueID: "1",
		UserID:  "vangel@elastecad.com",
		Ts:      time.Now(),
	}

	ar := []*entities.Resolution{r1, r2}

	js, err := json.Marshal(ar)

	assert.Nil(t, err)
	assert.NotEqual(t, js, "")
}
