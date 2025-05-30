package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SuratIzin struct {
	ID    int
	Jenis string
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
	PajakImpor   float64
	PajakEkspor  float64
	BeratBarang  float64
	SuratIzin    SuratIzin
	Transportasi Transportasi
}

var daftarBarang = [10]Barang{
	{
		Nama:        "Padi Organik",
		Jenis:       "Pertanian",
		Harga:       150000,
		PajakImpor:  0,
		PajakEkspor: 5,
		BeratBarang: 50,
		SuratIzin:   SuratIzin{ID: 1, Jenis: "Impor"},
		Transportasi: Transportasi{
			ID: "TR001", Jenis: "Truk", BiayaPerKM: 2000, WaktuPengiriman: 3,
		},
	},
	{
		Nama:        "Susu Sapi",
		Jenis:       "Peternakan",
		Harga:       100000,
		PajakImpor:  10,
		PajakEkspor: 0,
		BeratBarang: 20,
		SuratIzin:   SuratIzin{ID: 2, Jenis: "Ekspor"},
		Transportasi: Transportasi{
			ID: "TR002", Jenis: "Kereta", BiayaPerKM: 1500, WaktuPengiriman: 5,
		},
	},
	{
		Nama:        "Benih Jagung",
		Jenis:       "Pertanian",
		Harga:       50000,
		PajakImpor:  0,
		PajakEkspor: 0,
		BeratBarang: 10,
		SuratIzin:   SuratIzin{ID: 3, Jenis: "Lokal"},
		Transportasi: Transportasi{
			ID: "TR001", Jenis: "Truk", BiayaPerKM: 2000, WaktuPengiriman: 3,
		},
	},
	{
		Nama:        "Kopi Arabika",
		Jenis:       "Perkebunan",
		Harga:       250000,
		PajakImpor:  5,
		PajakEkspor: 8,
		BeratBarang: 15,
		SuratIzin:   SuratIzin{ID: 4, Jenis: "Ekspor"},
		Transportasi: Transportasi{
			ID: "TR003", Jenis: "Pesawat", BiayaPerKM: 5000, WaktuPengiriman: 1,
		},
	},
	{
		Nama:        "Coklat Fermentasi",
		Jenis:       "Perkebunan",
		Harga:       300000,
		PajakImpor:  15,
		PajakEkspor: 10,
		BeratBarang: 25,
		SuratIzin:   SuratIzin{ID: 5, Jenis: "Impor"},
		Transportasi: Transportasi{
			ID: "TR004", Jenis: "Kapal", BiayaPerKM: 1000, WaktuPengiriman: 7,
		},
	},
	{
		Nama:        "Daging Ayam",
		Jenis:       "Peternakan",
		Harga:       120000,
		PajakImpor:  8,
		PajakEkspor: 0,
		BeratBarang: 30,
		SuratIzin:   SuratIzin{ID: 6, Jenis: "Impor"},
		Transportasi: Transportasi{
			ID: "TR002", Jenis: "Kereta", BiayaPerKM: 1500, WaktuPengiriman: 5,
		},
	},
	{
		Nama:        "Ikan Laut Beku",
		Jenis:       "Perikanan",
		Harga:       200000,
		PajakImpor:  12,
		PajakEkspor: 5,
		BeratBarang: 40,
		SuratIzin:   SuratIzin{ID: 7, Jenis: "Ekspor"},
		Transportasi: Transportasi{
			ID: "TR003", Jenis: "Pesawat", BiayaPerKM: 5000, WaktuPengiriman: 1,
		},
	},
	{
		Nama:        "Kacang Tanah",
		Jenis:       "Pertanian",
		Harga:       80000,
		PajakImpor:  0,
		PajakEkspor: 4,
		BeratBarang: 12,
		SuratIzin:   SuratIzin{ID: 8, Jenis: "Lokal"},
		Transportasi: Transportasi{
			ID: "TR001", Jenis: "Truk", BiayaPerKM: 2000, WaktuPengiriman: 3,
		},
	},
	{
		Nama:        "Buah Naga",
		Jenis:       "Perkebunan",
		Harga:       180000,
		PajakImpor:  0,
		PajakEkspor: 6,
		BeratBarang: 22,
		SuratIzin:   SuratIzin{ID: 9, Jenis: "Ekspor"},
		Transportasi: Transportasi{
			ID: "TR004", Jenis: "Kapal", BiayaPerKM: 1000, WaktuPengiriman: 7,
		},
	},
	{
		Nama:        "Garam Laut",
		Jenis:       "Perikanan",
		Harga:       60000,
		PajakImpor:  5,
		PajakEkspor: 3,
		BeratBarang: 18,
		SuratIzin:   SuratIzin{ID: 10, Jenis: "Impor"},
		Transportasi: Transportasi{
			ID: "TR002", Jenis: "Kereta", BiayaPerKM: 1500, WaktuPengiriman: 5,
		},
	},
}

func cetakBarang(b Barang) {
	fmt.Println("Nama Barang       :", b.Nama)
	fmt.Println("Jenis             :", b.Jenis)
	fmt.Println("Harga             : Rp", b.Harga)
	fmt.Printf("Pajak Impor       : %.2f%%\n", b.PajakImpor)
	fmt.Printf("Pajak Ekspor      : %.2f%%\n", b.PajakEkspor)
	fmt.Println("Berat             :", b.BeratBarang, "kg")
	fmt.Println("Surat Izin        :", b.SuratIzin.ID, "-", b.SuratIzin.Jenis)
	fmt.Printf("Transportasi      : %s (Biaya per KM: Rp %d, Waktu: %d hari)\n", b.Transportasi.Jenis, b.Transportasi.BiayaPerKM, b.Transportasi.WaktuPengiriman)
	fmt.Println("-----------------------------------")
}

func cariBarang(nama string) int {
	nama = strings.ToLower(nama)
	for i := range daftarBarang {
		if strings.Contains(strings.ToLower(daftarBarang[i].Nama), nama) {
			return i
		}
	}
	return -1
}

func urutkanSuratIzinID() {
	n := len(daftarBarang)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if daftarBarang[i].SuratIzin.ID > daftarBarang[j].SuratIzin.ID {
				daftarBarang[i], daftarBarang[j] = daftarBarang[j], daftarBarang[i]
			}
		}
	}
}

func cariSuratIzinDenganBinary(id int) int {
	low, high := 0, len(daftarBarang)-1
	for low <= high {
		mid := (low + high) / 2
		if daftarBarang[mid].SuratIzin.ID == id {
			return mid
		} else if daftarBarang[mid].SuratIzin.ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func urutkanPajakImpor(asc bool) {
	n := len(daftarBarang)
	for i := 0; i < n-1; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if (asc && daftarBarang[j].PajakImpor < daftarBarang[index].PajakImpor) ||
				(!asc && daftarBarang[j].PajakImpor > daftarBarang[index].PajakImpor) {
				index = j
			}
		}
		daftarBarang[i], daftarBarang[index] = daftarBarang[index], daftarBarang[i]
	}
}

func urutkanHarga(asc bool) {
	n := len(daftarBarang)
	for i := 1; i < n; i++ {
		key := daftarBarang[i]
		j := i - 1
		for j >= 0 {
			lebih := daftarBarang[j].Harga > key.Harga
			if !asc {
				lebih = daftarBarang[j].Harga < key.Harga
			}
			if lebih {
				daftarBarang[j+1] = daftarBarang[j]
				j--
			} else {
				break
			}
		}
		daftarBarang[j+1] = key
	}
}

func urutkanNama() {
	n := len(daftarBarang)
	for i := 1; i < n; i++ {
		key := daftarBarang[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftarBarang[j].Nama) > strings.ToLower(key.Nama) {
			daftarBarang[j+1] = daftarBarang[j]
			j--
		}
		daftarBarang[j+1] = key
	}
}

func estimasiBiaya() {
	nama := input("Masukkan nama barang: ")
	index := cariBarang(nama)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}
	barang := daftarBarang[index]

	fmt.Println("Pilih jenis surat izin:")
	fmt.Println("1. Impor")
	fmt.Println("2. Ekspor")
	fmt.Print("Masukkan pilihan (1/2): ")

	var pilihan int
	fmt.Scanln(&pilihan)

	// Tentukan jenis surat izin berdasarkan pilihan
	switch pilihan {
	case 1:
		barang.SuratIzin.Jenis = "Impor"
	case 2:
		barang.SuratIzin.Jenis = "Ekspor"
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	jarak := 100.0
	biayaTransport := float64(barang.Transportasi.BiayaPerKM) * jarak
	pajak := 0.0

	switch barang.SuratIzin.Jenis {
	case "Impor":
		pajak = float64(barang.Harga) * barang.PajakImpor / 100
	case "Ekspor":
		pajak = float64(barang.Harga) * barang.PajakEkspor / 100
	}

	total := float64(barang.Harga) + pajak + biayaTransport

	fmt.Println("=== Estimasi Biaya", barang.SuratIzin.Jenis, "===")
	fmt.Printf("Harga Barang       : Rp %.0f\n", float64(barang.Harga))
	fmt.Printf("Pajak %s        : Rp %.0f\n", barang.SuratIzin.Jenis, pajak)
	fmt.Printf("Biaya Transportasi : Rp %.0f\n", biayaTransport)
	fmt.Printf("Total Biaya        : Rp %.0f\n", total)
}

func tampilkanSemuaBarang() {
	for _, barang := range daftarBarang {
		if barang.Nama != "" {
			cetakBarang(barang)
		}
	}
}

func input(teks string) string {
	fmt.Print(teks)
	reader := bufio.NewReader(os.Stdin)
	hasil, _ := reader.ReadString('\n')
	return strings.TrimSpace(hasil)
}

func menu() {
	for {
		fmt.Println("\n=== Sistem Logistik ===")
		fmt.Println("1. Cari Barang (Sequential Search)")
		fmt.Println("2. Cari Surat Izin (Binary Search)")
		fmt.Println("3. Estimasi Biaya Impor/Ekspor")
		fmt.Println("4. Tampilkan Semua Barang")
		fmt.Println("5. Urutkan Pajak Impor (Selection Sort)")
		fmt.Println("6. Urutkan Harga Barang (Insertion Sort)")
		fmt.Println("7. Keluar")

		pilihan := input("Pilih menu (1-7): ")

		switch pilihan {
		case "1":
			nama := input("Masukkan nama barang: ")
			index := cariBarang(nama)
			if index == -1 {
				fmt.Println("Barang tidak ditemukan.")
			} else {
				cetakBarang(daftarBarang[index])
			}
		case "2":
			urutkanNama()
			urutkanSuratIzinID()
			var id int
			fmt.Print("Masukkan ID Surat Izin yang dicari: ")
			fmt.Scanln(&id)
			index := cariSuratIzinDenganBinary(id)
			if index == -1 {
				fmt.Println("Barang dengan ID Surat Izin tersebut tidak ditemukan.")
			} else {
				cetakBarang(daftarBarang[index])
			}
		case "3":
			estimasiBiaya()
		case "4":
			tampilkanSemuaBarang()
		case "5":
			urut := input("Urutkan pajak impor secara menaik? (y/n): ")
			urutkanPajakImpor(strings.ToLower(urut) == "y")
			fmt.Println("Data telah diurutkan berdasarkan Pajak Impor.")
			tampilkanSemuaBarang()
		case "6":
			urut := input("Urutkan harga secara menaik? (y/n): ")
			urutkanHarga(strings.ToLower(urut) == "y")
			fmt.Println("Data telah diurutkan berdasarkan Harga.")
			tampilkanSemuaBarang()
		case "7":
			fmt.Println("Terima kasih telah menggunakan sistem.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func main() {
	menu()
}
