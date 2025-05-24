# Sistem-Informasi-Logistik-Ekspor-dan-Impor


📦 Sistem Informasi Logistik Ekspor dan Impor
Sistem Informasi Logistik Ekspor dan Impor adalah sebuah aplikasi berbasis Command Line Interface (CLI) yang dikembangkan menggunakan bahasa pemrograman Golang. Aplikasi ini dirancang untuk membantu pengelolaan logistik pada aktivitas ekspor dan impor di Indonesia, dengan fitur-fitur analisis dan pencarian berbasis data dari file JSON.

Sistem ini memungkinkan pengguna untuk mencari, memfilter, menyortir, dan menghitung estimasi biaya barang secara cepat dan efisien, mendukung pengambilan keputusan dalam rantai pasok logistik secara lebih tepat dan terstruktur.

🧩 Fitur Utama
🔍 Pencarian Nama Barang
Cari data barang berdasarkan nama.

📄 Pencarian Surat Perizinan
Telusuri barang berdasarkan jenis surat izin.

🚚 Penyortiran Transportasi Barang
Filter barang berdasarkan jenis transportasi seperti laut, udara, atau darat.

📦 Penyortiran Jenis Barang
Tampilkan barang sesuai kategori jenisnya (misal: elektronik, tekstil, dsb.).

💰 Estimasi Biaya Logistik
Hitung estimasi biaya berdasarkan harga barang, pajak, dan jarak tempuh dengan formula:

Total = Harga + (Harga × Pajak%) + (Jarak × Tarif/km)

📊 Tampilan Seluruh Data Barang
Lihat seluruh data logistik dalam satu tampilan.

🔢 Penyortiran Data

Heap Sort: berdasarkan persentase pajak.

Radix Sort (simulasi): berdasarkan harga barang.

📋 Menu Interaktif
Navigasi berbasis angka memudahkan pengguna dalam memilih fitur.

❌ Exit Program
Keluar dari aplikasi dengan aman.

⚙️ Teknologi
Bahasa Pemrograman: Go (Golang)

Input Data: File JSON (data.json)

Eksekusi Program: Command Line Interface (CLI)

📌 Tujuan
Aplikasi ini ditujukan untuk membantu pengusaha, pengelola logistik, dan pihak terkait dalam melakukan pencatatan, analisis, dan estimasi biaya logistik ekspor-impor secara efisien. Sistem ini juga merupakan bagian dari Tugas Besar (TuBes) yang menggambarkan penerapan nyata pengembangan sistem informasi berbasis data dan proses logistik terotomasi.
