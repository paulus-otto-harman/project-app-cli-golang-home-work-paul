package views

import (
	"homework/component"
	"homework/util"
)

func Home() {
	for {
		util.ClearScreen()
		println("*** MENU UTAMA ***")
		println("[1] Tambah Buku")
		println("[2] Hapus Buku")
		println("[3] Tampilkan Buku")
		println("[4] Keluar dari Program")

		menuItem, err := component.ToInt(component.Input(component.Args(component.P("type", "number"), component.P("label", "Masukkan Pilihan Anda :"))))
		switch {
		case err != nil:
			continue
		case menuItem == 4:
			return
		case menuItem == 3:
			TampilkanBuku()
		case menuItem == 2:
			HapusBuku()
		case menuItem == 1:
			TambahBuku()
		}
	}
}
