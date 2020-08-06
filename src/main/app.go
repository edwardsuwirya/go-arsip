package main

import (
	"github.com/rivo/tview"
)

var app *tview.Application
var pages *tview.Pages
var daftarArsip = make([]arsip, 0, 1)

func main() {
	//var arsip01 = arsip{
	//	nomor:      "001",
	//	judul:      "Arsip A",
	//	keterangan: "Arsip A",
	//}
	//var arsip02 = arsip{
	//	nomor:      "002",
	//	judul:      "Arsip B",
	//	keterangan: "Arsip B",
	//}
	//
	//var arsip03 = arsip{
	//	nomor:      "003",
	//	judul:      "Arsip C",
	//	keterangan: "Arsip C",
	//}
	//tambahArsip(arsip01)
	//tambahArsip(arsip02)
	//tambahArsip(arsip03)
	initUi()
}

func initUi() {
	app = tview.NewApplication()

	pages = tview.NewPages().
		AddPage("halamanUtama", halamanUtama(), true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
