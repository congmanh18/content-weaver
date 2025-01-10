package entity

type EbookIndex struct {
	ID        string `json:"id"`
	EbookID   int    `json:"ebook_id"`
	IndexName string `json:"index_name"`
}

func (e *EbookIndex) TableName() string {
	return "pdf.ebook_indexes"
}
