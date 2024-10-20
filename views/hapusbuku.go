package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
)

func HapusBuku(rak models.Rak) {
	util.ClearScreen()
	fmt.Println("*** Hapus Buku ***")
	isbn, _ := component.ToString(component.Input(map[string]interface{}{"label": "Isbn"}))
	rak.RemoveBook(isbn)
}
