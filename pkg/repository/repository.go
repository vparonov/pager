package repository

type Repository interface {
	UpsertIssueType(typeName string, template string) error
	FindIssueType(typeName string) (string, bool)
}
