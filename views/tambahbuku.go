package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
)

func simpan(rak models.Perpustakaan, buku models.Buku) {
	rak.AddBook(&buku)
}

func TambahBuku(rak models.Rak) {

	for {
		util.ClearScreen()

		fmt.Println("*** Tambah Buku ***")

		judul, _ := component.ToString(component.Input(map[string]interface{}{"label": "Judul"}))
		penulis, _ := component.ToString(component.Input(map[string]interface{}{"label": "Penulis"}))
		isbn, _ := component.ToString(component.Input(map[string]interface{}{"label": "Isbn"}))

		simpan(&rak, models.InitBuku(judul, penulis, isbn))

		menuItem, err := component.Input(map[string]interface{}{"type": "number", "label": "Tekan Enter untuk menambahkan lagi atau Masukkan [0] untuk Kembali Ke Menu Utama"})
		switch {
		case err != nil:
			continue
		case menuItem == 0:
			return
		}

	}

}
