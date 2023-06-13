package repository

import (
	"template-go/internal/collection"
)

type Repository interface {
	CollectionRepository() collection.Repository
}
