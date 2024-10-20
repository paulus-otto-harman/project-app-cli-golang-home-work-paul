package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
)

func TambahBuku() {

loop:
	for {
		util.ClearScreen()

		fmt.Println("*** Tambah Buku ***")

		judul, _ := component.ToString(component.Input(map[string]interface{}{"label": "Judul"}))
		penulis, _ := component.ToString(component.Input(map[string]interface{}{"label": "Penulis"}))
		isbn, _ := component.ToString(component.Input(map[string]interface{}{"label": "Isbn"}))

		models.InitBuku(judul, penulis, isbn)

		menuItem, err := component.Input(map[string]interface{}{"type": "number", "label": "Tekan Enter untuk menambahkan lagi atau Masukkan [0] untuk Kembali Ke Menu Utama"})
		fmt.Println(menuItem, err)
		switch {
		case err != nil:
			continue
		case menuItem == 0:
			break loop
		}

	}

}
