package main

import (
	"github.com/rivo/tview"
)

//
//import (
//	ui "github.com/gizak/termui/v3"
//	"github.com/gizak/termui/v3/widgets"
//)
//

//
type arsip struct {
	nomor      string
	judul      string
	keterangan string
}

func halamanTambahArsip() (f *tview.Form) {

	var noArsip, judulArsip, keterangan string
	f = tview.NewForm().
		AddInputField("No Arsip", "", 20, nil, func(text string) {
			noArsip = text
		}).
		AddInputField("Judul Arsip", "", 20, nil, func(text string) {
			judulArsip = text
		}).
		AddInputField("Keterangan", "", 20, nil, func(text string) {
			keterangan = text
		}).
		AddButton("Simpan", func() {
			daftarArsip = append(daftarArsip, arsip{
				nomor:      noArsip,
				judul:      judulArsip,
				keterangan: keterangan,
			})
			pages.AddAndSwitchToPage("Success", successNotif(), true)
		}).
		AddButton("Batal", func() {
			pages.SwitchToPage("halamanUtama")
		})
	f.SetBorder(true).SetTitle("Tambah Arsip").SetTitleAlign(tview.AlignLeft)
	return
}

func halamanListArsip() (flex *tview.Flex) {
	t := tview.NewTable().SetBorders(true)
	t.SetBorder(true)
	t.SetCell(0, 0, tview.NewTableCell("No Arsip"))
	t.SetCell(0, 1, tview.NewTableCell("Judul Arsip"))
	t.SetCell(0, 2, tview.NewTableCell("Keterangan"))

	for r := 1; r <= len(daftarArsip); r++ {
		t.SetCell(r, 0,
			tview.NewTableCell(daftarArsip[r-1].nomor))
		t.SetCell(r, 1,
			tview.NewTableCell(daftarArsip[r-1].judul))
		t.SetCell(r, 2,
			tview.NewTableCell(daftarArsip[r-1].keterangan))
	}

	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(navigasi("Daftar Arsip"), 0, 1, true).
		AddItem(t, 0, 4, false)

	return
}
