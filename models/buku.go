package models

type Buku struct {
	judul, penulis, isbn string
}

func InitBuku(judul string, penulis string, isbn string) Buku {
	return Buku{judul, penulis, isbn}
}

func (buku *Buku) Judul() string {
	return buku.judul
}

func (buku *Buku) Penulis() string {
	return buku.penulis
}

func (buku *Buku) Isbn() string {
	return buku.isbn
}
