package views

import (
	"fmt"
	"homework/component"
	"homework/models"
	"homework/util"
	"strings"
)

func TampilkanBuku() {
	const WidthJudul = 60
	const WidthPenulis = 60
	const WidthIsbn = 30

	util.ClearScreen()

	// header
	fmt.Printf("%s%s%s%s%s%s%s\n", "┌─", strings.Repeat("─", WidthJudul), "┬─", strings.Repeat("─", WidthPenulis), "┬─", strings.Repeat("─", WidthIsbn), "┐")
	fmt.Printf("%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthJudul, "Judul", "│ ", WidthPenulis, "Penulis", "│ ", WidthIsbn, "ISBN", "│")
	fmt.Printf("%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthJudul), "┼─", strings.Repeat("─", WidthPenulis), "┼─", strings.Repeat("─", WidthIsbn), "┤")

	// body
	for _, buku := range models.DaftarBuku {
		fmt.Printf("%s%-*s%s%-*s%s%-*s%s\n", "│ ", WidthJudul, buku.Judul(), "│ ", WidthPenulis, buku.Penulis(), "│ ", WidthIsbn, buku.Isbn(), "│")
	}

	// footer
	fmt.Printf("%s%s%s%s%s%s%s\n", "├─", strings.Repeat("─", WidthJudul), "┴─", strings.Repeat("─", WidthPenulis), "┴─", strings.Repeat("─", WidthIsbn), "┤")
	fmt.Printf("%s%-*s%-*s%-*s%s\n", "│ ", WidthJudul, fmt.Sprintf("%s%d%s", "Total : ", len(models.DaftarBuku), " buku"), WidthPenulis+2, "", WidthIsbn+2, "", "│")
	fmt.Printf("%s%s%s\n", "╘═", strings.Repeat("═", WidthJudul+2+WidthPenulis+2+WidthIsbn), "╛")

	component.Input(map[string]interface{}{"label": "Tekan Enter untuk kembali ke Menu Utama ..."})
}
