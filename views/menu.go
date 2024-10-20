package views

import (
	"homework/component"
	"homework/controllers"
	"homework/util"
)

func Home() {

	var menuItem interface{}
	var err error

	for {
		util.ClearScreen()
		println("*** MENU UTAMA ***")
		println("[1] Tambah Buku")
		println("[2] Hapus Buku")
		println("[3] Tampilkan Buku")
		println("[4] Keluar dari Program")

		menuItem, err = component.Input(map[string]interface{}{"type": "number", "label": "Masukkan Pilihan Anda :"})
		switch {
		case err != nil:
			continue
		case menuItem == 4:
			break
		case menuItem == 3:
			controllers.TampilkanBuku()
		case menuItem == 2:
			controllers.HapusBuku()
		case menuItem == 1:
			controllers.TambahBuku()
		}
	}
}
