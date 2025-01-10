package application

import "github.com/congmanh18/content-weaver/ports"

type PDFService struct {
	Processor ports.PDFProcessor
	Repo      ports.EbookRepository
}

func (s *PDFService) ProcessAndSaveEbook(filePath string) error {
	return nil
}
