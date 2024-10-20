package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
)

func HapusBuku(rak models.Rak) {
	for {
		util.ClearScreen()
		fmt.Println("*** Hapus Buku ***")
		isbn, _ := component.ToString(component.Input(map[string]interface{}{"label": fmt.Sprintf("%-15s :", "ISBN")}))
		err := rak.RemoveBook(isbn)
		if err == nil {
			fmt.Println("Buku berhasil dihapus!")

		} else {
			fmt.Println(err)
		}
		menuItem, err := component.Input(component.Args(component.P("type", "number"), component.P("label", fmt.Sprintf("\n\n%s", "Tekan Enter untuk menghapus yang lain atau Masukkan [0] untuk Kembali Ke Menu Utama"))))
		switch {
		case err != nil:
			continue
		case menuItem == 0:
			return
		}
	}

}
