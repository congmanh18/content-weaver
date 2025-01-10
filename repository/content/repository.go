package content

import (
	"github.com/congmanh18/content-weaver/core/db/postgresql"
	"github.com/congmanh18/content-weaver/entity"
)

type Repo interface {
	SaveContent(ebook []entity.EbookContent) error
}

type contentImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &contentImpl{DB: db}
}
