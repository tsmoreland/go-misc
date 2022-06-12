package shared

type RepositoryFactory interface {
	Build() (Repository, error)
}
