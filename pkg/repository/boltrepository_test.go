package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoltRepository(t *testing.T) {
	os.Remove("test.db")

	repo := NewBoltRepository("test.db")

	err := repo.Open()

	assert.Nil(t, err)

	err = repo.UpsertIssueType("test", "test")

	assert.Nil(t, err)

	err = repo.Close()

	assert.Nil(t, err)
}
