package models

var DaftarBuku = []Buku{}

type Buku struct {
	judul, penulis, isbn string
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

func (buku *Buku) IsiJudul(judul string) string {
	buku.judul = judul
	return buku.judul
}

func (buku *Buku) IsiPenulis(penulis string) string {
	buku.penulis = penulis
	return buku.penulis
}

func (buku *Buku) IsiIsbn(isbn string) string {
	buku.isbn = isbn
	return buku.isbn
}

func (buku *Buku) Tambah() string {
	DaftarBuku = append(DaftarBuku, *buku)
	return "Buku Berhasil Ditambah"
}

func (buku *Buku) CariIsbn(isbn string) int {
	for i, buku := range DaftarBuku {
		if isbn == buku.isbn {
			return i
		}
	}
	return -1
}

func (buku *Buku) Hapus(index int) string {
	DaftarBuku = append(DaftarBuku[:index], DaftarBuku[index+1:]...)
	return "Buku Berhasil Dihapus"
}
