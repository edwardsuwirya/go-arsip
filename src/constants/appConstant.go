package constants

//Page ID Constants
const (
	MainPageId         = "halamanUtama"
	AddDocPageId       = "halamanTambahArsip"
	ViewListDocPageId  = "halamanLihatArsip"
	SuccessNotifPageId = "notifikasiSukses"
	ErrorNotifPageId   = "notifikasiGagal"
)

//Message type
const (
	SuccessMessageType = iota
	FailedMessageType
)

//UI Label
var AppName = map[string]string{"EN": "Document Management System", "ID": "Sistem Informasi Manajemen Arsip"}
var HomeLabel = map[string]string{"EN": "Home", "ID": "Beranda"}
var ExitLabel = map[string]string{"EN": "Exit", "ID": "Keluar"}
var AddDocLabel = map[string]string{"EN": "Add Document", "ID": "Tambah Arsip"}
var ViewDocLabel = map[string]string{"EN": "View Document", "ID": "Lihat Arsip"}

var AddDoc2ndLabel = map[string]string{"EN": "Add your new document", "ID": "Menambahkan arsip baru"}
var ViewDoc2ndLabel = map[string]string{"EN": "View your document list", "ID": "Lihat semua arsip"}
var Exit2ndLabel = map[string]string{"EN": "Exit application", "ID": "Keluar aplikasi"}

var SuccessLabel = map[string]string{"EN": "Success", "ID": "Sukses"}
var FailedLabel = map[string]string{"EN": "Failed", "ID": "Gagal"}
var MandatoryFailedLabel = map[string]string{"EN": "Some field is required", "ID": "Lengkapi data"}

var DocNoLabel = map[string]string{"EN": "Document #", "ID": "No. Arsip"}
var DocTitleLabel = map[string]string{"EN": "Title", "ID": "Judul Arsip"}
var NotesLabel = map[string]string{"EN": "Notes", "ID": "Keterangan"}

var DocTotalLabel = map[string]string{"EN": "Total documents", "ID": "Total arsip"}

var SaveLabel = map[string]string{"EN": "Save", "ID": "Simpan"}
var CancelLabel = map[string]string{"EN": "Cancel", "ID": "Batal"}

var AddMoreLabel = map[string]string{"EN": "Create Another", "ID": "Buat Lagi"}
var OkLabel = map[string]string{"EN": "Ok", "ID": "Ya"}
var FinishLabel = map[string]string{"EN": "Finish", "ID": "Selesai"}
