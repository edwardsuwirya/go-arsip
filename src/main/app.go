package main

import (
	"github.com/rivo/tview"
	"os"
	"strings"
)

var app *tview.Application
var pages *tview.Pages
var daftarArsip = make([]Arsip, 0, 1)
var defLang = "ID"
var fileName = "arsip.csv"

func main() {
	if len(os.Args) == 2 {
		langArgs := os.Args[1]
		if len(langArgs) > 1 {
			defLangTemp := strings.ToUpper(langArgs)
			if defLang == "ID" || defLang == "EN" {
				defLang = defLangTemp
			}
		}
	}
	ReadCsv(fileName)
	initUi()
}

func initUi() {
	app = tview.NewApplication()

	pages = tview.NewPages().
		AddPage(MainPageId, halamanUtama(), true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
