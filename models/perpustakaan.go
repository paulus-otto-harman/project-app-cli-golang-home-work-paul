package models

import (
	"errors"
	"fmt"
	"homework/component"
	"homework/util"
	"strings"
)

type Perpustakaan interface {
	AddBook(buku *Buku)
	RemoveBook(isbn string) error
	ShowBooks()
}

var daftarBuku = []Buku{}

type Rak struct {
}

func (r *Rak) AddBook(buku *Buku) {
	daftarBuku = append(daftarBuku, *buku)
}

func cariIsbn(isbn string) int {
	for i, buku := range daftarBuku {
		if isbn == buku.isbn {
			return i
		}
	}
	return -1
}

func (rak *Rak) RemoveBook(isbn string) error {
	index := cariIsbn(isbn)
	if index == -1 {
		return errors.New(fmt.Sprintf("Buku dengan ISBN %s tidak ditemukan", isbn))
	}
	daftarBuku = append(daftarBuku[:index], daftarBuku[index+1:]...)
	return nil
}

func (rak *Rak) ShowBooks() {
	const WidthJudul = 60
	const WidthPenulis = 60
	const WidthIsbn = 30

	util.ClearScreen()

	// header
	fmt.Printf("%s%s%s%s%s%s%s\n", "┌─", strings.Repeat("─", WidthJudul), "┬─", strings.Repeat("─", WidthPenulis), "┬─", strings.Repeat("─", WidthIsbn), "┐")
	fmt.Printf("%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthJudul, "Judul", "│ ", WidthPenulis, "Penulis", "│ ", WidthIsbn, "ISBN", "│")
	fmt.Printf("%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthJudul), "┼─", strings.Repeat("─", WidthPenulis), "┼─", strings.Repeat("─", WidthIsbn), "┤")

	// body
	for _, buku := range daftarBuku {
		fmt.Printf("%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthJudul, buku.Judul(), "│ ", WidthPenulis, buku.Penulis(), "│ ", WidthIsbn, buku.Isbn(), "│")
	}

	// footer
	fmt.Printf("%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthJudul), "┴─", strings.Repeat("─", WidthPenulis), "┴─", strings.Repeat("─", WidthIsbn), "┤")
	fmt.Printf("%s%-*s%-*s%-*s%s\n", "│ ", WidthJudul, fmt.Sprintf("%s%d%s", "Total : ", len(daftarBuku), " buku"), WidthPenulis+2, "", WidthIsbn+2, "", "│")
	fmt.Printf("%s%s%s\n", "╘═", strings.Repeat("═", WidthJudul+2+WidthPenulis+2+WidthIsbn), "╛")

	component.Input(map[string]interface{}{"label": "Tekan Enter untuk kembali ke Menu Utama ..."})
}
