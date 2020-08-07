package main

import (
	"github.com/rivo/tview"
	"strconv"
)

func halamanUtama() (flex *tview.Flex) {
	list := tview.NewList().
		AddItem(AddDocLabel[defLang], AddDoc2ndLabel[defLang], 'd', func() {
			gantiHalaman(AddDocPageId, halamanTambahArsip())
		}).
		AddItem(ViewDocLabel[defLang], ViewDoc2ndLabel[defLang], 'l', func() {
			gantiHalaman(ViewListDocPageId, halamanListArsip(daftarArsip))
		}).
		AddItem(ExitLabel[defLang], Exit2ndLabel[defLang], 'q', func() {
			app.Stop()
		})
	list.SetBorder(true).SetTitle(AppName[defLang])

	t := tview.NewTable()
	t.SetBorder(true).SetTitle("Dashboard")
	t.SetCell(0, 0, tview.NewTableCell(DocTotalLabel[defLang]))
	t.SetCell(0, 1, tview.NewTableCell(":"))
	t.SetCell(0, 2, tview.NewTableCell(strconv.Itoa(len(daftarArsip))))
	flex = tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(t, 0, 4, false)
	return
}

func gantiHalaman(idHalaman string, komponen tview.Primitive) {
	pages.RemovePage(idHalaman)
	pages.AddAndSwitchToPage(idHalaman, komponen, true)
}

func navigasi(title string) (nav *tview.List) {
	nav = tview.NewList().
		AddItem(HomeLabel[defLang], "", 'h', func() {
			pages.SwitchToPage(MainPageId)
		})
	nav.SetBorder(true).SetTitle(title)
	return
}

func showMessage(message string, buttons []string, callback func(int, string)) (modal *tview.Modal) {
	modal = tview.NewModal().
		SetText(message).
		AddButtons(buttons).
		SetDoneFunc(callback)
	return
}
