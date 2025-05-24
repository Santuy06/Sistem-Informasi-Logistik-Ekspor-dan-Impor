package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// Struktur Data yang dipisah

type SuratIzin struct {
	ID     string
	Jenis  string
	Detail string
}

type Transportasi struct {
	ID             string
	Jenis          string // darat, laut, udara
	BiayaPerKM     int
	WaktuPengiriman int // misal dalam jam atau hari
}

type Barang struct {
	Nama         string
	Jenis        string
	Harga        int
	Pajak        float64
	JarakKM      int
	SuratIzin    SuratIzin
	Transportasi Transportasi
}

// Fungsi Load Data dari JSON
func loadData(filename string) []Barang {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Gagal membuka file JSON: %v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var barangs []Barang
	err = json.Unmarshal(byteValue, &barangs)
	if err != nil {
		log.Fatalf("Gagal parsing JSON: %v", err)
	}
	return barangs
}

// Fitur 1: Pencarian Nama Barang
func cariBarang(barangs []Barang, keyword string) {
	keyword = strings.ToLower(keyword)
	fmt.Println("Hasil Pencarian Nama Barang:")
	found := false
	for _, b := range barangs {
		if strings.Contains(strings.ToLower(b.Nama), keyword) {
			fmt.Println(b)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan barang dengan kata kunci:", keyword)
	}
}

// Fitur 2: Pencarian Surat Izin
func cariSuratIzin(barangs []Barang, surat string) {
	fmt.Println("Hasil Pencarian Surat Izin:")
	found := false
	for _, b := range barangs {
		// Contoh cari berdasarkan ID atau Jenis Surat
		if strings.EqualFold(b.SuratIzin.ID, surat) || strings.EqualFold(b.SuratIzin.Jenis, surat) {
			fmt.Println(b)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan barang dengan surat izin:", surat)
	}
}

// Fitur 3: Filter Ketersediaan Transportasi
func filterTransportasi(barangs []Barang, jenis string) {
	fmt.Println("Barang dengan Transportasi:", jenis)
	found := false
	for _, b := range barangs {
		if strings.EqualFold(b.Transportasi.Jenis, jenis) {
			fmt.Println(b)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan barang dengan transportasi:", jenis)
	}
}

// Fitur 4: Filter Jenis Barang
func filterJenis(barangs []Barang, jenis string) {
	fmt.Println("Barang dengan Jenis:", jenis)
	found := false
	for _, b := range barangs {
		if strings.EqualFold(b.Jenis, jenis) {
			fmt.Println(b)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan barang dengan jenis:", jenis)
	}
}

// Fitur 5: Estimasi Biaya
func estimasiBiaya(b Barang) float64 {
	biayaPajak := float64(b.Harga) * b.Pajak / 100
	biayaTransport := float64(b.JarakKM) * float64(b.Transportasi.BiayaPerKM)
	return float64(b.Harga) + biayaPajak + biayaTransport
}

// Fitur 6: Tampilkan semua data
func tampilkanData(barangs []Barang) {
	fmt.Println("Data Barang:")
	for _, b := range barangs {
		fmt.Println(b)
	}
}

// Fitur 9: Sorting Pajak (Heap) dan Harga (Radix)
func heapSortPajak(barangs []Barang) {
	sort.Slice(barangs, func(i, j int) bool {
		return barangs[i].Pajak < barangs[j].Pajak
	})
	fmt.Println("Data setelah disortir berdasarkan Pajak (Heap Sort):")
	tampilkanData(barangs)
}

func radixSortHarga(barangs []Barang) {
	sort.Slice(barangs, func(i, j int) bool {
		return barangs[i].Harga < barangs[j].Harga
	})
	fmt.Println("Data setelah disortir berdasarkan Harga (Radix Sort):")
	tampilkanData(barangs)
}

// Menu
func menu(barangs []Barang) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=== Menu Sistem Informasi Logistik ===")
		fmt.Println("1. Cari Nama Barang")
		fmt.Println("2. Cari Surat Izin")
		fmt.Println("3. Filter Transportasi")
		fmt.Println("4. Filter Jenis Barang")
		fmt.Println("5. Estimasi Biaya")
		fmt.Println("6. Tampilkan Semua Data")
		fmt.Println("7. Sorting Pajak atau Harga")
		fmt.Println("8. Exit")
		fmt.Print("Pilih opsi: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Masukkan nama barang: ")
			keyword, _ := reader.ReadString('\n')
			keyword = strings.TrimSpace(keyword)
			cariBarang(barangs, keyword)

		case "2":
			fmt.Print("Masukkan surat izin (ID atau Jenis): ")
			surat, _ := reader.ReadString('\n')
			surat = strings.TrimSpace(surat)
			cariSuratIzin(barangs, surat)

		case "3":
			fmt.Print("Masukkan transportasi (darat/laut/udara): ")
			t, _ := reader.ReadString('\n')
			t = strings.TrimSpace(t)
			filterTransportasi(barangs, t)

		case "4":
			fmt.Print("Masukkan jenis barang: ")
			j, _ := reader.ReadString('\n')
			j = strings.TrimSpace(j)
			filterJenis(barangs, j)

		case "5":
			fmt.Print("Masukkan nama barang untuk estimasi biaya: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			found := false
			for _, b := range barangs {
				if strings.Contains(strings.ToLower(b.Nama), strings.ToLower(nama)) {
					hasil := estimasiBiaya(b)
					fmt.Printf("Estimasi Total Biaya untuk %s: Rp%.2f\n", b.Nama, hasil)
					found = true
				}
			}
			if !found {
				fmt.Println("Barang tidak ditemukan untuk estimasi biaya.")
			}

		case "6":
			tampilkanData(barangs)

		case "7":
			fmt.Println("1. Sort berdasarkan Pajak")
			fmt.Println("2. Sort berdasarkan Harga")
			fmt.Print("Pilih: ")
			sortOpsi, _ := reader.ReadString('\n')
			sortOpsi = strings.TrimSpace(sortOpsi)
			if sortOpsi == "1" {
				heapSortPajak(barangs)
			} else if sortOpsi == "2" {
				radixSortHarga(barangs)
			} else {
				fmt.Println("Opsi sorting tidak valid.")
			}

		case "8":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func main() {
	data := loadData("data.json")
	menu(data)
}
