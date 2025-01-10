package ports

import "github.com/congmanh18/content-weaver/entity"

type PDFProcessor interface {
	ExtractMetadata(filePath *string) (entity.Ebook, error)
	ExtractContent(filePath *string) ([]entity.EbookContent, error)
}
