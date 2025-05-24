package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Barang struct {
	Nama     string  `json:"nama"`
	Kategori string  `json:"kategori"`
	Harga    float64 `json:"harga"`
	Pajak    float64 `json:"pajak"`
}

type SuratIzin struct {
	ID       string `json:"id"`
	Jenis    string `json:"jenis"`
	Deskripsi string `json:"deskripsi"`
}

type Transportasi struct {
	Jenis         string  `json:"jenis"`
	WaktuKirim    int     `json:"waktu_kirim"`
	Biaya         float64 `json:"biaya"`
	JarakPengirim float64 `json:"jarak_pengiriman"`
}

var daftarBarang []Barang
var daftarIzin []SuratIzin
var daftarTransportasi []Transportasi

func main() {
	loadData()
	menu()
}

func loadData() {
	// Data dummy
	daftarBarang = []Barang{
		{"Laptop", "elektronik", 12000000, 1500000},
		{"Kain Batik", "tekstil", 50000, 5000},
		{"Beras", "bahan makanan", 10000, 1000},
		{"TV", "elektronik", 3000000, 300000},
		{"Kemeja", "tekstil", 150000, 15000},
	}

	daftarIzin = []SuratIzin{
		{"IZ001", "Ekspor", "Izin untuk ekspor barang"},
		{"IZ002", "Impor", "Izin untuk impor barang"},
	}

	daftarTransportasi = []Transportasi{
		{"darat", 3, 200000, 100},
		{"laut", 7, 500000, 500},
		{"udara", 1, 1000000, 1000},
	}
}

func menu() {
	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Pencarian Nama Barang")
		fmt.Println("2. Pencarian Surat Perizinan")
		fmt.Println("3. Penyortiran Transportasi Barang")
		fmt.Println("4. Filter Jenis Barang")
		fmt.Println("5. Estimasi Biaya")
		fmt.Println("6. Tampilan Data")
		fmt.Println("7. Sort Harga dan Pajak")
		fmt.Println("8. Simpan ke File JSON")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu (1-9): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			searchBarang()
		case 2:
			searchIzin()
		case 3:
			sortTransportasi()
		case 4:
			filterJenisBarang()
		case 5:
			estimasiBiaya()
		case 6:
			displayData()
		case 7:
			sortHargaDanPajak()
		case 8:
			saveToFileJSON()
		case 9:
			fmt.Println("Keluar dari program...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func searchBarang() {
	fmt.Print("Masukkan nama barang: ")
	var nama string
	fmt.Scan(&nama)
	for _, b := range daftarBarang {
		if strings.Contains(strings.ToLower(b.Nama), strings.ToLower(nama)) {
			fmt.Printf("Nama: %s | Kategori: %s | Harga: %.2f\n", b.Nama, b.Kategori, b.Harga)
		}
	}
}

func searchIzin() {
	fmt.Print("Masukkan ID atau jenis surat izin: ")
	var keyword string
	fmt.Scan(&keyword)
	for _, i := range daftarIzin {
		if strings.Contains(strings.ToLower(i.ID), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(i.Jenis), strings.ToLower(keyword)) {
			fmt.Printf("ID: %s | Jenis: %s | Deskripsi: %s\n", i.ID, i.Jenis, i.Deskripsi)
		}
	}
}

func sortTransportasi() {
	sort.Slice(daftarTransportasi, func(i, j int) bool {
		return daftarTransportasi[i].Biaya < daftarTransportasi[j].Biaya
	})
	fmt.Println("Data transportasi disortir berdasarkan biaya:")
	for _, t := range daftarTransportasi {
		fmt.Printf("Jenis: %s | Waktu Kirim: %d hari | Biaya: %.2f\n", t.Jenis, t.WaktuKirim, t.Biaya)
	}
}

func filterJenisBarang() {
	fmt.Print("Masukkan kategori barang (contoh: tekstil): ")
	var kategori string
	fmt.Scan(&kategori)
	for _, b := range daftarBarang {
		if strings.ToLower(b.Kategori) == strings.ToLower(kategori) {
			fmt.Printf("Nama: %s | Kategori: %s | Harga: %.2f\n", b.Nama, b.Kategori, b.Harga)
		}
	}
}

func estimasiBiaya() {
	fmt.Print("Masukkan nama barang: ")
	var nama string
	fmt.Scan(&nama)
	var barang *Barang
	for i := range daftarBarang {
		if strings.ToLower(daftarBarang[i].Nama) == strings.ToLower(nama) {
			barang = &daftarBarang[i]
			break
		}
	}
	if barang == nil {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan jenis transportasi: ")
	var jenis string
	fmt.Scan(&jenis)
	var transport *Transportasi
	for i := range daftarTransportasi {
		if strings.ToLower(daftarTransportasi[i].Jenis) == strings.ToLower(jenis) {
			transport = &daftarTransportasi[i]
			break
		}
	}
	if transport == nil {
		fmt.Println("Transportasi tidak ditemukan.")
		return
	}
	total := barang.Harga + transport.Biaya + barang.Pajak
	fmt.Printf("Estimasi biaya total: %.2f\n", total)
}

func displayData() {
	fmt.Println("\n=== Daftar Barang ===")
	for _, b := range daftarBarang {
		fmt.Printf("Nama: %s | Kategori: %s | Harga: %.2f | Pajak: %.2f\n", b.Nama, b.Kategori, b.Harga, b.Pajak)
	}
	fmt.Println("\n=== Daftar Surat Izin ===")
	for _, i := range daftarIzin {
		fmt.Printf("ID: %s | Jenis: %s | Deskripsi: %s\n", i.ID, i.Jenis, i.Deskripsi)
	}
	fmt.Println("\n=== Daftar Transportasi ===")
	for _, t := range daftarTransportasi {
		fmt.Printf("Jenis: %s | Waktu Kirim: %d hari | Biaya: %.2f | Jarak: %.2f\n", t.Jenis, t.WaktuKirim, t.Biaya, t.JarakPengirim)
	}
}

func sortHargaDanPajak() {
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Harga")
	fmt.Println("2. Pajak")
	var opsi int
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		sort.Slice(daftarBarang, func(i, j int) bool {
			return daftarBarang[i].Harga < daftarBarang[j].Harga
		})
	case 2:
		sort.Slice(daftarBarang, func(i, j int) bool {
			return daftarBarang[i].Pajak < daftarBarang[j].Pajak
		})
	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}
	for _, b := range daftarBarang {
		fmt.Printf("Nama: %s | Harga: %.2f | Pajak: %.2f\n", b.Nama, b.Harga, b.Pajak)
	}
}

func saveToFileJSON() {
	data := map[string]interface{}{
		"barang":       daftarBarang,
		"surat_izin":   daftarIzin,
		"transportasi": daftarTransportasi,
	}

	file, err := os.Create("data_export.json")
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Gagal menulis ke file:", err)
		return
	}

	fmt.Println("Data berhasil disimpan ke data_export.json")
}
