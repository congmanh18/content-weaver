package outline

import (
	"github.com/congmanh18/content-weaver/core/db/postgresql"
	"github.com/congmanh18/content-weaver/entity"
)

type Repo interface {
	SaveOutline(ebook []entity.EbookOutline) error
}

type outlineImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &outlineImpl{DB: db}
}
