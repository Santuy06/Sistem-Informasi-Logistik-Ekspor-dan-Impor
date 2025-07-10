# Sistem-Informasi-Logistik-Ekspor-dan-Impor

# Estimasi Biaya Pengiriman Barang

Program ini ditulis menggunakan bahasa Go (Golang) untuk menghitung estimasi biaya pengiriman suatu barang berdasarkan nama barang, jenis surat izin (Impor, Ekspor, atau Lokal), serta jenis transportasi.

## Fitur

- Menyimpan daftar barang dengan informasi:
  - Nama dan jenis barang.
  - Harga barang dan berat.
  - Jenis surat izin (Impor, Ekspor, Lokal).
  - Pajak impor dan ekspor.
  - Data transportasi (jenis, biaya per KM, waktu pengiriman).
- Memungkinkan pengguna memilih barang berdasarkan nama.
- Menghitung total biaya transportasi.
- Menghitung pajak berdasarkan jenis izin.
- Menampilkan hasil estimasi biaya secara jelas.

## Struktur Data

Program ini menggunakan beberapa `struct`, seperti:

```go
type Barang struct {
    Nama         string
    Jenis        string
    Harga        float64
    PajakImpor   float64
    PajakEkspor  float64
    BeratBarang  float64
    SuratIzin    SuratIzin
    Transportasi Transportasi
}

type SuratIzin struct {
    ID    int
    Jenis string // Impor, Ekspor, atau Lokal
}

type Transportasi struct {
    ID              string
    Jenis           string // Truk, Kereta, Kapal, Pesawat
    BiayaPerKM      float64
    WaktuPengiriman int // dalam hari
}
