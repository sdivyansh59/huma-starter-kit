package greeting

import (
	"github.com/uptrace/bun"
)

type IRepository interface {
	// Define methods that the repository should implement
}

type Repository struct {
	db *bun.DB
}

func NewRepository() *Repository {
	return &Repository{
		// Initialize any fields if necessary
	}
}

var _ IRepository = (*Repository)(nil)
