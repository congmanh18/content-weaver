package entity

type Ebook struct {
	ID        string
	Title     string
	Author    string
	Subject   string
	PageCount int
	FileLink  string
}

func (e *Ebook) TableName() string {
	return "pdf.ebooks"
}
