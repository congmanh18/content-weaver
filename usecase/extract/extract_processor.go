package extract

import (
	"github.com/congmanh18/content-weaver/entity"
)

type ExtractPDFProcessor interface {
	ExtractMetadata(filePath string) (entity.Ebook, error)
	ExtractOutline(filePath string) ([]entity.EbookOutline, error)
	ExtractContent(filePath string) ([]entity.EbookContent, error)
}
