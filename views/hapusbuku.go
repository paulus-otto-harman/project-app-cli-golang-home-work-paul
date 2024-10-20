package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
)

func HapusBuku() {
	util.ClearScreen()
	buku := models.Buku{}

	fmt.Println("*** Hapus Buku ***")
	isbn, _ := component.Input(map[string]interface{}{"label": "Isbn"})
	buku.Hapus(buku.CariIsbn(fmt.Sprintf("%v", isbn)))
}
