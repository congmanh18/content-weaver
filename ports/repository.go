package ports

import "github.com/congmanh18/content-weaver/entity"

type EbookRepository interface {
	SaveEbook(ebook entity.Ebook) error
	GetEbooks() ([]entity.Ebook, error)
}
