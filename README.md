# Sistem-Informasi-Logistik-Ekspor-dan-Impor

---

# 📦 Sistem Informasi Logistik Ekspor dan Impor

**Sistem Informasi Logistik Ekspor dan Impor** adalah aplikasi berbasis **Command Line Interface (CLI)** yang dikembangkan menggunakan **bahasa pemrograman Go (Golang)**. Aplikasi ini dirancang untuk mendukung pengelolaan data logistik pada aktivitas ekspor dan impor di Indonesia secara cepat, efisien, dan terstruktur.

Dengan memanfaatkan file data berformat **JSON**, aplikasi ini menyediakan berbagai fitur pencarian, penyortiran, serta estimasi biaya logistik untuk mendukung pengambilan keputusan dalam rantai pasok.

---

## 🧩 Fitur Utama

### 🔍 Pencarian Nama Barang

Cari dan temukan informasi barang berdasarkan nama.

### 📄 Pencarian Berdasarkan Surat Izin

Filter data barang berdasarkan jenis **Surat Izin** seperti *Impor*, *Ekspor*, atau *Lokal*.

### 🚚 Penyortiran Transportasi

Tampilkan barang berdasarkan jenis transportasi: **Truk**, **Kereta**, **Pesawat**, atau **Kapal**.

### 📦 Penyortiran Jenis Barang

Kelompokkan barang sesuai kategori seperti: **Pertanian**, **Peternakan**, **Perikanan**, **Perkebunan**, dll.

### 💰 Estimasi Biaya Logistik

Hitung estimasi total biaya logistik dengan rumus:

```
Total Biaya = Harga + (Harga × Pajak%) + (Jarak × Tarif per KM)
```

### 📊 Tampilan Seluruh Data

Lihat seluruh data barang secara lengkap dalam satu tampilan CLI.

### 🔢 Penyortiran Data

* **Heap Sort** → Menyortir berdasarkan persentase pajak (tinggi ke rendah).
* **Radix Sort (Simulasi)** → Menyortir berdasarkan harga barang.

### 📋 Menu Interaktif

Navigasi mudah berbasis angka untuk mengakses berbagai fitur.

### ❌ Exit Program

Keluar dari aplikasi dengan aman dan tertib.

---

## ⚙️ Teknologi

* **Bahasa Pemrograman**: Go (Golang)
* **Input Data**: File JSON (`data.json`)
* **Antarmuka**: Command Line Interface (CLI)

---

## 📌 Tujuan

Aplikasi ini dikembangkan sebagai bagian dari **Tugas Besar (TuBes)** untuk menggambarkan penerapan nyata dari pengembangan **sistem informasi berbasis data** dalam konteks **logistik ekspor dan impor**.

Diharapkan dapat membantu:

* Pengusaha
* Pengelola logistik
* Pemerintah dan pemangku kepentingan lain

Dalam melakukan **pencatatan, analisis, dan estimasi biaya** logistik secara lebih efisien dan terotomatisasi.

---
