package repository

import (
	"errors"

	"github.com/vparonov/pager/pkg/entities"
	bolt "go.etcd.io/bbolt"
)

type boltRepository struct {
	fileName string
	db       *bolt.DB
}

func NewBoltRepository(fileName string) Repository {
	return &boltRepository{
		fileName: fileName,
	}
}

func (repo *boltRepository) Open() error {
	var err error

	repo.db, err = bolt.Open(repo.fileName, 0600, nil)
	return err
}

func (repo *boltRepository) Close() error {
	var err error
	if repo.db != nil {
		err = repo.db.Close()
		repo.db = nil
	}
	return err
}
func (r *boltRepository) UpsertIssueType(typeName string, template string) error {
	return r.db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte("template"))

		if err != nil {
			return err
		}

		return bucket.Put([]byte(typeName), []byte(template))
	})
}

func (r *boltRepository) FindIssueType(typeName string) (string, bool) {
	var rawtemplate []byte
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("template"))

		if bucket == nil {
			return errors.New("bucket 'template' not found")
		}

		// nil byte slice converts to empty string
		// so there is no need to check for nil here
		rawtemplate = bucket.Get([]byte(typeName))
		return nil
	})

	if err != nil {
		return "", false
	}

	if rawtemplate == nil {
		return "", false
	}

	return string(rawtemplate), true
}

func (r *boltRepository) InsertIssue(issue *entities.Issue) error {
	return r.db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte("issue"))

		if err != nil {
			return err
		}

		return bucket.Put([]byte(issue.ID), []byte(issue.Body))
	})
}

func (r *boltRepository) FindIssue(id string) (*entities.Issue, error) {
	return nil, nil
}
