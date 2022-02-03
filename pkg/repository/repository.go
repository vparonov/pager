package repository

// NB! All functions should be thread safe
type Repository interface {
	UpsertIssueType(typeName string, template string) error
	FindIssueType(typeName string) (string, bool)

	InsertIssue(id string, body string) error
	
}
