package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

type Arsip struct {
	Nomor      string
	Judul      string
	Keterangan string
}

var inputField *tview.InputField

func halamanTambahArsip() (flex *tview.Flex) {

	var noArsip, judulArsip, keterangan string
	f := tview.NewForm().
		AddInputField(DocNoLabel[defLang], "", 20, nil, func(text string) {
			noArsip = text
		}).
		AddInputField(DocTitleLabel[defLang], "", 20, nil, func(text string) {
			judulArsip = text
		}).
		AddInputField(NotesLabel[defLang], "", 20, nil, func(text string) {
			keterangan = text
		}).
		AddButton(SaveLabel[defLang], func() {
			daftarArsip = append(daftarArsip, Arsip{
				Nomor:      noArsip,
				Judul:      judulArsip,
				Keterangan: keterangan,
			})
			CreateCsv(fileName, daftarArsip)
			gantiHalaman(SuccessNotifPageId, successNotif())
		}).
		AddButton(CancelLabel[defLang], func() {
			gantiHalaman(MainPageId, halamanUtama())
		})
	f.SetBorder(true).SetTitle(AddDocLabel[defLang]).SetTitleAlign(tview.AlignLeft)

	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(f, 0, 4, true)
	return
}

func halamanListArsip(list []Arsip) (flex *tview.Flex) {
	t := tview.NewTable().SetBorders(true)
	t.SetBorder(true).SetTitle("Daftar Arsip")
	t.SetCell(0, 0, tview.NewTableCell(DocNoLabel[defLang]))
	t.SetCell(0, 1, tview.NewTableCell(DocTitleLabel[defLang]))
	t.SetCell(0, 2, tview.NewTableCell(NotesLabel[defLang]))

	for r := 1; r <= len(list); r++ {
		t.SetCell(r, 0,
			tview.NewTableCell(list[r-1].Nomor))
		t.SetCell(r, 1,
			tview.NewTableCell(list[r-1].Judul))
		t.SetCell(r, 2,
			tview.NewTableCell(list[r-1].Keterangan))
	}
	nav := navigasi("")

	inputField = tview.NewInputField().
		SetLabel("Enter a number: ").
		SetPlaceholder("E.g. 1234").
		SetFieldWidth(20).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				cariArsipDenganNomor()
			}
			if key == tcell.KeyTAB {
				app.SetFocus(nav)
			}
		})
	inputField.SetBorder(true).SetTitle("Cari Arsip").SetTitleAlign(tview.AlignLeft)
	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 0, 1, true).
		AddItem(t, 0, 4, false).
		AddItem(nav, 0, 1, true)

	return
}

func cariArsipDenganNomor() {
	hasil := make([]Arsip, 0, 1)
	pencarian := inputField.GetText()
	for _, a := range daftarArsip {
		if strings.HasPrefix(a.Nomor, pencarian) {
			hasil = append(hasil, a)
		}
	}
	gantiHalaman(ViewListDocPageId, halamanListArsip(hasil))
}
func successNotif() (modal *tview.Modal) {
	return showMessage(SuccessLabel[defLang], []string{"Ok"}, func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Ok" {
			pages.RemovePage(SuccessNotifPageId)
			gantiHalaman(AddDocPageId, halamanTambahArsip())
		}
	})
}
