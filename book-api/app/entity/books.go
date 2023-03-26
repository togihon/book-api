package entity

type Book struct {
	BookID      string `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}
