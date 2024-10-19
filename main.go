package main

import (
	"homework/coms"
	"homework/controllers"
	"homework/views"
)

func xmain() {
	views.TambahBuku()
}

func main() {
	var menuItem interface{}
	var err error
loop:
	for {
		views.HomeMenu()
		menuItem, err = coms.Input(map[string]interface{}{"type": "number", "label": "Masukkan Pilihan Anda :"})
		switch {
		case err != nil:
			continue
		case menuItem == 4:
			break loop
		case menuItem == 3:
			controllers.TampilkanBuku()
		case menuItem == 2:
			controllers.HapusBuku()
		case menuItem == 1:
			controllers.TambahBuku()
		}
	}
}
