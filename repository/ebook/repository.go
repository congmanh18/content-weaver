package ebook

import (
	"github.com/congmanh18/content-weaver/core/db/postgresql"
	"github.com/congmanh18/content-weaver/entity"
)

type Repo interface {
	SaveEbook(ebook entity.Ebook) error
}

type ebookImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &ebookImpl{DB: db}
}
