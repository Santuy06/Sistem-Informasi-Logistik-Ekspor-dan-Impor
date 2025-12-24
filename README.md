# Sistem Ekspor Impor 🚢

Sistem manajemen ekspor dan impor barang yang dibangun dengan arsitektur modern (Frontend-Backend terpisah) dan dirancang untuk menangani beban data besar (500.000+ records).

## 🚀 Teknologi yang Digunakan

- **Frontend**: React.js (Vite), Ant Design (UI), Axios, Recharts (Visualisasi).
- **Backend**: Go (Gin Gonic), MySQL (Driver), GORM-like Raw SQL, JWT Auth, Bcrypt.
- **Testing**: JMeter (Performance/Stress Test), Locust (Python-based Load Testing).

## ✨ Fitur Utama

- **Manajemen Barang**: CRUD (Create, Read, Update, Delete) data barang dengan efisiensi tinggi.
- **Master Data**: Kelola data Petugas (dengan struktur hirarki/tree), Transportasi, dan Surat Izin.
- **Kalkulator Estimasi**: Penghitungan pajak dan biaya pengiriman berdasarkan jenis barang dan rute.
- **Autentikasi Keamanan**: Login aman dengan JWT (JSON Web Token) dan hashing password Bcrypt.
- **Keamanan API**: Implementasi rate limiting dan security headers (XSS Protection, CSP, dll).
- **Skalabilitas**: Database seed script untuk menghasilkan 500.000 data test secara otomatis.

## 🛠️ Persiapan & Instalasi

### 1. Database Setup
1. Pastikan MySQL sudah berjalan.
2. Impor skema dan jalankan seeding dari file `backend/database_setup.sql`:
   ```sql
   SOURCE backend/database_setup.sql;
   ```
   *Script ini akan otomatis menginisialisasi database `ekspor_impor` dan mengisi 500.000 data.*

### 2. Backend (Go)
1. Masuk ke folder `backend`.
2. Salin `.env.example` ke `.env` dan sesuaikan konfigurasi database.
3. Jalankan server:
   ```bash
   go run main.go
   ```
   *Backend akan berjalan di `http://localhost:8080`.*

### 3. Frontend (React)
1. Jalankan instalasi dependensi (di root folder):
   ```bash
   npm install
   ```
2. Jalankan development server:
   ```bash
   npm run dev
   ```
   *Frontend akan berjalan di `http://localhost:5173`.*

## 📈 Pengujian Performa (Performance Testing)

Proyek ini menyertakan file pengujian untuk memastikan sistem tetap stabil pada beban tinggi:
- **JMeter**: Gunakan `dashboard_test.jmx` atau `sql_perf_test.jmx`.
- **Locust**: Jalankan `locustfile.py` untuk mensimulasikan trafik pengguna secara konkuren.

## 📁 Struktur Folder

- `/src` - Source code Frontend React.
- `/backend` - Source code Backend Go & SQL Scripts.
- `/database` - File penunjang database tambahan.
- `*.jmx` & `locustfile.py` - Script pengujian performa.

---
*Dibuat untuk keperluan sistem manajemen logistik yang tangguh dan terukur.*
