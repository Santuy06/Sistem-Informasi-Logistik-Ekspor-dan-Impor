// kode program

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Struktur Data
type SuratIzin struct {
	ID     string
	Jenis  string
	Detail string
}

type Transportasi struct {
	ID              string
	Jenis           string
	BiayaPerKM      int
	WaktuPengiriman int
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

// Fungsi baca data dari JSON
func muatData(namaFile string) []Barang {
	file, err := os.Open(namaFile)
	if err != nil {
		log.Fatalf("Gagal membuka file JSON: %v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var daftarBarang []Barang
	err = json.Unmarshal(byteValue, &daftarBarang)
	if err != nil {
		log.Fatalf("Gagal parsing JSON: %v", err)
	}
	return daftarBarang
}

// Fitur pencarian nama barang
func cariNamaBarang(daftar []Barang, kataKunci string) {
	kataKunci = strings.ToLower(kataKunci)
	ditemukan := false
	fmt.Println("Hasil pencarian:")
	for _, barang := range daftar {
		if strings.Contains(strings.ToLower(barang.Nama), kataKunci) {
			fmt.Println(barang)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ditemukan barang dengan kata kunci tersebut.")
	}
}

// Fitur pencarian surat izin
func cariSurat(daftar []Barang, surat string) {
	ditemukan := false
	fmt.Println("Hasil pencarian surat izin:")
	for _, barang := range daftar {
		if strings.EqualFold(barang.SuratIzin.ID, surat) || strings.EqualFold(barang.SuratIzin.Jenis, surat) {
			fmt.Println(barang)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ditemukan surat izin tersebut.")
	}
}

// Filter berdasarkan transportasi
func filterTransportasi(daftar []Barang, jenis string) {
	fmt.Println("Barang dengan transportasi:", jenis)
	ada := false
	for _, barang := range daftar {
		if strings.EqualFold(barang.Transportasi.Jenis, jenis) {
			fmt.Println(barang)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ditemukan barang dengan jenis transportasi tersebut.")
	}
}

// Filter berdasarkan jenis barang
func filterJenisBarang(daftar []Barang, jenis string) {
	fmt.Println("Barang dengan jenis:", jenis)
	ada := false
	for _, barang := range daftar {
		if strings.EqualFold(barang.Jenis, jenis) {
			fmt.Println(barang)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ditemukan barang dengan jenis tersebut.")
	}
}

// Estimasi biaya
func estimasiBiaya(barang Barang) float64 {
	biayaPajak := float64(barang.Harga) * barang.Pajak / 100
	biayaTransport := float64(barang.JarakKM) * float64(barang.Transportasi.BiayaPerKM)
	return float64(barang.Harga) + biayaPajak + biayaTransport
}

// Tampilkan semua data
func tampilkanSemua(daftar []Barang) {
	for _, barang := range daftar {
		fmt.Println(barang)
	}
}

// Manual Heap Sort berdasarkan Pajak
func heapify(arr []Barang, n, i int) {
	terbesar := i
	kiri := 2*i + 1
	kanan := 2*i + 2

	if kiri < n && arr[kiri].Pajak > arr[terbesar].Pajak {
		terbesar = kiri
	}
	if kanan < n && arr[kanan].Pajak > arr[terbesar].Pajak {
		terbesar = kanan
	}
	if terbesar != i {
		arr[i], arr[terbesar] = arr[terbesar], arr[i]
		heapify(arr, n, terbesar)
	}
}

func heapSortPajak(arr []Barang) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
	fmt.Println("Data disortir berdasarkan Pajak:")
	tampilkanSemua(arr)
}

func MaksHarga(arr []Barang) int {
	maks := arr[0].Harga
	for _, barang := range arr {
		if barang.Harga > maks {
			maks = barang.Harga
		}
	}
	return maks
}

func HitungSortHarga(arr []Barang, exp int) {
	n := len(arr)
	output := make([]Barang, n)
	hitung := make([]int, 10)

	for i := 0; i < n; i++ {
		idx := (arr[i].Harga / exp) % 10
		hitung[idx]++
	}
	for i := 1; i < 10; i++ {
		hitung[i] += hitung[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		idx := (arr[i].Harga / exp) % 10
		output[hitung[idx]-1] = arr[i]
		hitung[idx]--
	}
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// Manual Radix Sort berdasarkan Pajak
func radixSortHarga(arr []Barang) {
	maks := MaksHarga(arr)
	for exp := 1; maks/exp > 0; exp *= 10 {
		HitungSortHarga(arr, exp)
	}
	fmt.Println("Data disortir berdasarkan Harga:")
	tampilkanSemua(arr)
}

// Menu utama
func menu(daftar []Barang) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=== Menu Sistem Logistik ===")
		fmt.Println("1. Cari Nama Barang")
		fmt.Println("2. Cari Surat Izin")
		fmt.Println("3. Filter Transportasi")
		fmt.Println("4. Filter Jenis Barang")
		fmt.Println("5. Estimasi Biaya")
		fmt.Println("6. Tampilkan Semua Data")
		fmt.Println("7. Sorting Pajak atau Harga")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")

		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			fmt.Print("Nama barang: ")
			nama, _ := reader.ReadString('\n')
			cariNamaBarang(daftar, strings.TrimSpace(nama))
		case "2":
			fmt.Print("Surat izin (ID/Jenis): ")
			surat, _ := reader.ReadString('\n')
			cariSurat(daftar, strings.TrimSpace(surat))
		case "3":
			fmt.Print("Jenis transportasi: ")
			tr, _ := reader.ReadString('\n')
			filterTransportasi(daftar, strings.TrimSpace(tr))
		case "4":
			fmt.Print("Jenis barang: ")
			jenis, _ := reader.ReadString('\n')
			filterJenisBarang(daftar, strings.TrimSpace(jenis))
		case "5":
			fmt.Print("Nama barang: ")
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			ditemukan := false
			for _, barang := range daftar {
				if strings.Contains(strings.ToLower(barang.Nama), strings.ToLower(nama)) {
					fmt.Printf("Estimasi total biaya %s: Rp%.2f\n", barang.Nama, estimasiBiaya(barang))
					ditemukan = true
				}
			}
			if !ditemukan {
				fmt.Println("Barang tidak ditemukan.")
			}
		case "6":
			tampilkanSemua(daftar)
		case "7":
			fmt.Println("1. Pajak")
			fmt.Println("2. Harga")
			fmt.Print("Pilih: ")
			sortOpsi, _ := reader.ReadString('\n')
			sortOpsi = strings.TrimSpace(sortOpsi)
			if sortOpsi == "1" {
				heapSortPajak(daftar)
			} else if sortOpsi == "2" {
				radixSortHarga(daftar)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		case "8":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Menu tidak dikenali.")
		}
	}
}

func main() {
	data := muatData("data.json")
	menu(data)
}
