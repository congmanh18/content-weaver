package entity

type EbookContent struct {
	ID      string `json:"id"`
	EbookID int    `json:"ebook_id"`
	IndexID int    `json:"index_id"`
	Content string `json:"content"`
}

func (e *EbookContent) TableName() string {
	return "pdf.ebook_content"
}
