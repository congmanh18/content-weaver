package usecase

import (
	"github.com/congmanh18/content-weaver/models"
	"github.com/congmanh18/content-weaver/repository/content"
	"github.com/congmanh18/content-weaver/repository/ebook"
	"github.com/congmanh18/content-weaver/repository/outline"
	"github.com/congmanh18/content-weaver/usecase/extract"
)

type Usecase interface {
	ProcessAndSaveEbook(filePath string) models.Error
}

type PDFService struct {
	Processor extract.ExtractPDFProcessor
	ebook     ebook.Repo
	outline   outline.Repo
	content   content.Repo
}

func (s *PDFService) ProcessAndSaveEbook(filePath string) *models.Error {
	// Bước 1: Trích xuất metadata
	metadata, err := s.Processor.ExtractMetadata(filePath)
	if err != nil {
		return &models.Error{
			Code:    1001,
			Message: "Error extracting metadata",
			Err:     err,
		}
	}

	// Bước 2: Trích xuất mục lục (outline)
	outline, err := s.Processor.ExtractOutline(filePath)
	if err != nil {
		return &models.Error{
			Code:    1002,
			Message: "Error extracting outline",
			Err:     err,
		}
	}

	// Bước 3: Trích xuất nội dung
	contents, err := s.Processor.ExtractContent(filePath)
	if err != nil {
		return &models.Error{
			Code:    1003,
			Message: "Error extracting content",
			Err:     err,
		}
	}

	// Bước 4: Lưu ebook vào repository
	err = s.ebook.SaveEbook(metadata)
	if err != nil {
		return &models.Error{
			Code:    1004,
			Message: "Error saving ebook to repository",
			Err:     err,
		}
	}

	// Bước 5: Lưu mục lục vào repository
	err = s.outline.SaveOutline(outline)
	if err != nil {
		return &models.Error{
			Code:    1005,
			Message: "Error saving outline to repository",
			Err:     err,
		}
	}

	// Bước 6: Lưu nội dung vào repository
	err = s.content.SaveContent(contents)
	if err != nil {
		return &models.Error{
			Code:    1006,
			Message: "Error saving content to repository",
			Err:     err,
		}
	}
	// Thành công
	return nil
}
