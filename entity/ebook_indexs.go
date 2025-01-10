package entity

type EbookOutline struct {
	ID          string `json:"id"`
	EbookID     int    `json:"ebook_id"`
	OutlineName string `json:"outline_name"`
}

func (e *EbookOutline) TableName() string {
	return "pdf.ebook_outlines"
}
