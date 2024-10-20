package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
)

func Home() {
	rak := models.Rak{}
	for {
		util.ClearScreen()
		fmt.Println("*** MANAJEMEN PERPUSTAKAAN ***")
		fmt.Println("[1] Tambah Buku")
		fmt.Println("[2] Hapus Buku")
		fmt.Println("[3] Tampilkan Buku")
		fmt.Println("[4] Keluar dari Program")

		menuItem, err := component.ToInt(component.Input(component.Args(component.P("type", "number"), component.P("label", "Masukkan Pilihan Anda :"))))
		switch {
		case err != nil:
			continue
		case menuItem == 4:
			return
		case menuItem == 3:
			//TampilkanBuku(rak)
			rak.ShowBooks()
		case menuItem == 2:
			HapusBuku(rak)
		case menuItem == 1:
			TambahBuku(rak)
		}
	}
}
