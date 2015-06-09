package cms

type ResourceRepo interface {
	Store(resource Resource)
	Delete(resource Resource)
	FindByPath(path string)
}

type Resource struct {
	Path       string
	Title      string
	ShortTitle string
}

type SqlResourceRepo struct {
	address string
}

func NewSqlResourceRepo(address string) *SqlResourceRepo {
	repo := new(SqlResourceRepo)
	repo.address = address
	return repo
}
