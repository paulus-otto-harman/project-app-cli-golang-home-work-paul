package models

type Perpustakaan interface {
	AddBook(buku *Buku)
	RemoveBook(isbn string) error
	ShowBooks()
}
