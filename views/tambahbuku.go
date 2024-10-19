package views

import (
	"fmt"
	"homework/coms"
	"homework/models"
	"homework/utils"
)

func TambahBuku() {

loop:
	for {
		utils.ClearScreen()
		fmt.Println("*** Tambah Buku ***")
		judul, _ := coms.Input(map[string]interface{}{"label": "Judul"})
		penulis, _ := coms.Input(map[string]interface{}{"label": "Penulis"})
		isbn, _ := coms.Input(map[string]interface{}{"label": "Isbn"})

		buku := models.Buku{}
		buku.IsiJudul(fmt.Sprintf("%v", judul))
		buku.IsiPenulis(fmt.Sprintf("%v", penulis))
		buku.IsiIsbn(fmt.Sprintf("%v", isbn))
		buku.Tambah()

		menuItem, err := coms.Input(map[string]interface{}{"type": "number", "label": "Tekan Enter untuk menambahkan lagi atau Masukkan [0] untuk Kembali Ke Menu Utama"})
		fmt.Println(menuItem, err)
		switch {
		case err != nil:
			continue
		case menuItem == 0:
			break loop
		}

	}

}
