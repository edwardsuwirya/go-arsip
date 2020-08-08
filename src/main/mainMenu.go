package main

import (
	"fmt"
	"github.com/gdamore/tcell"
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

func navigasi(judul string) (nav *tview.List) {
	nav = tview.NewList().
		AddItem(HomeLabel[defLang], "", 'h', func() {
			pages.SwitchToPage(MainPageId)
		})
	nav.SetBorder(true).SetTitle(judul)
	return
}

func notifikasi(pesan string, jenisPesan int, tombol []string, callback func(int, string)) (modal *tview.Modal) {
	color := tcell.ColorGreen
	if jenisPesan == FailedMessageType {
		color = tcell.ColorRed
	}
	modal = tview.NewModal().
		SetBackgroundColor(color).
		SetText(pesan).
		AddButtons(tombol).
		SetDoneFunc(callback)
	return
}

func labelDibutuhkan(label string) string {
	return fmt.Sprintf("%s %s", label, "*")
}
