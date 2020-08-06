package main

import "github.com/rivo/tview"

func halamanUtama() (list *tview.List) {
	list = tview.NewList().
		AddItem("Tambah Arsip", "Menambahkan Arsip Baru", 'd', func() {
			pages.AddAndSwitchToPage("halamanTambahArsip", halamanTambahArsip(), true)
		}).
		AddItem("Lihat Arsip", "Lihat Semua Arsip", 'l', func() {
			pages.AddAndSwitchToPage("halamanLihatArsip", halamanListArsip(), true)
		}).
		AddItem("Keluar", "Keluar Aplikasi", 'q', func() {
			app.Stop()
		})
	list.SetBorder(true).SetTitle("Sistem Informasi Manajemen Arsip")
	return
}

func navigasi(title string) (nav *tview.List) {
	nav = tview.NewList().
		AddItem("Beranda", "", 'h', func() {
			pages.SwitchToPage("halamanUtama")
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

func successNotif() (modal *tview.Modal) {
	return showMessage("Success", []string{"Ok"}, func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Ok" {
			pages.SwitchToPage("halamanUtama")
		}
	})

}
