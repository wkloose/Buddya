# Buddya Backend API

Buddya adalah layanan backend untuk aplikasi pemandu wisata dan budaya digital di Indonesia. Backend ini menyajikan informasi budaya daerah, rekomendasi destinasi wisata, dan tempat kuliner tradisional dari database lokal, sekaligus memanfaatkan kecerdasan buatan Google Gemini untuk mencari jadwal acara budaya terkini secara langsung.

---

## Teknologi yang Digunakan

- Bahasa Pemrograman: Go (Golang) versi 1.24
- Framework Web: Gin
- Database ORM: GORM
- Database: PostgreSQL
- Integrasi AI: Google Gemini API (model gemini-2.0-flash)
- Manajemen Konfigurasi: godotenv

---

## Fitur Utama

1. Informasi Budaya Kota: Menyajikan deskripsi dan ciri khas budaya dari berbagai kota di Indonesia.
2. Tempat Wisata dan Kuliner Lokal: Menyediakan data tempat wisata budaya dan kuliner khas lengkap dengan alamat serta titik koordinat lokasi.
3. Rekomendasi Acara Budaya Otomatis: Memanfaatkan Gemini API untuk mencari informasi acara atau festival budaya terkini di kota yang dipilih.
4. Pencatatan Riwayat AI (Prompt Logging): Menyimpan format prompt dan riwayat respons AI ke database untuk kemudahan peninjauan dan dokumentasi.

---

## Struktur Folder Proyek

- controllers: Menangani request HTTP yang masuk dan mengirimkan response JSON kembali ke pengguna.
- initializers: Mengatur konfigurasi awal aplikasi, seperti membaca file environment dan menghubungkan database.
- models: Berisi definisi struktur tabel database (City, Place, Event, PromptLog, PromptTemplate, User).
- Routes: Menentukan daftar alamat rute atau endpoint API.
- Seeder: Mengisi data awal ke database secara otomatis (data kota, tempat wisata, kuliner, dan template prompt).
- services: Berisi logika utama aplikasi, pengambilan data dari database, serta komunikasi dengan Gemini API.
- views dan public: Berkas tampilan antarmuka dan file statis bawaan.

---

## Prasyarat Sistem

Sebelum menjalankan aplikasi ini, pastikan komputer Anda sudah terpasang:

- Go versi 1.24 atau versi yang kompatibel
- PostgreSQL yang sedang aktif
- Google Gemini API Key

---

## Panduan Menjalankan Aplikasi

1. Buka Direktori Proyek
Buka terminal dan pastikan Anda berada di dalam folder proyek Buddya.

2. Konfigurasi File Lingkungan (.env)
Pastikan berkas .env sudah tersedia di folder utama proyek dan sesuaikan nilainya dengan konfigurasi komputer Anda:

```env
PORT=8080
DB_URL="host=localhost user=postgres password=password_anda dbname=buddya_db port=5432 sslmode=disable"
SECRET=kunci_rahasia_jwt
GEMINI_API_KEY=kunci_api_gemini_anda
```

3. Unduh Dependensi
Jalankan perintah berikut untuk mengunduh semua modul dan paket yang diperlukan:

```bash
go mod download
```

4. Jalankan Aplikasi
Jalankan server aplikasi dengan perintah:

```bash
go run main.go
```

Aplikasi akan berjalan pada port yang sudah diatur di file .env (misalnya http://localhost:8080).

---

## Daftar Endpoint API

Semua endpoint dapat diakses melalui awalan rute /api.

| Metode | Endpoint | Parameter Query | Deskripsi |
|---|---|---|---|
| GET | /api/budaya | city_name atau city_id | Menampilkan deskripsi budaya khas suatu kota |
| GET | /api/wisata | city_name atau city_id | Menampilkan daftar tempat wisata budaya di suatu kota |
| GET | /api/kuliner | city_name atau city_id | Menampilkan daftar kuliner khas di suatu kota |
| GET | /api/places | city_name dan category | Menampilkan tempat berdasarkan kota dan kategori (wisata atau kuliner) |
| GET | /api/events | city_name | Mencari jadwal acara budaya terkini di suatu kota melalui Gemini AI |
| GET | /api/all | city_name | Mengambil data lengkap (budaya, wisata, kuliner, dan acara budaya) sekaligus |

---

## Contoh Penggunaan API

### 1. Mengambil Informasi Lengkap Suatu Kota

Permintaan (Request):
```http
GET /api/all?city_name=Malang HTTP/1.1
Host: localhost:8080
```

Contoh Respons (Response):
```json
{
  "city": "Malang",
  "results": {
    "budaya": "Malang memiliki kebudayaan yang kental dengan kesenian Topeng Malangan dan tradisi peninggalan kerajaan.",
    "wisata": [
      {
        "id": 1,
        "nama_tempat": "Candi Singosari",
        "kategori": "wisata",
        "alamat": "Jl. Kertanegara No.148, Candirenggo, Singosari, Malang",
        "latitude": -7.892,
        "longitude": 112.664,
        "deskripsi": "Situs bersejarah peninggalan Kerajaan Singasari.",
        "image_url": "",
        "city_id": 1
      }
    ],
    "kuliner": [
      {
        "id": 2,
        "nama_tempat": "Bakso President",
        "kategori": "kuliner",
        "alamat": "Jl. Batanghari No.5, Rampal Celaket, Klojen, Malang",
        "latitude": -7.965,
        "longitude": 112.636,
        "deskripsi": "Kuliner bakso legendaris khas kota Malang.",
        "image_url": "",
        "city_id": 1
      }
    ],
    "events": [
      {
        "name": "Festival Topeng Malangan",
        "date": "2026-12-15",
        "location": "Balai Kota Malang",
        "description": "Pentas seni tari topeng dan pameran kerajinan khas Malang.",
        "link": "https://maps.google.com/..."
      }
    ]
  }
}
```

### 2. Mengambil Rekomendasi Acara Budaya Saja

Permintaan (Request):
```http
GET /api/events?city_name=Yogyakarta HTTP/1.1
Host: localhost:8080
```

Contoh Respons (Response):
```json
{
  "city_name": "Yogyakarta",
  "events": [
    {
      "name": "Pentas Wayang Kulit",
      "date": "2026-11-25",
      "location": "Kraton Ngayogyakarta Hadiningrat",
      "link": "https://maps.google.com/..."
    }
  ]
}
```

---

## Catatan Penting untuk Pengembangan

1. Pengaturan Ulang Database Otomatis:
Secara bawaan, file main.go menjalankan fungsi ResetDatabase() pada proses inisialisasi. Artinya, setiap kali server dimatikan dan dinyalakan kembali, tabel-tabel akan dihapus dan diisi ulang dari data seeder. Jika Anda ingin data log atau data baru tidak terhapus saat server dihidupkan ulang, Anda dapat mengganti fungsi tersebut menjadi SyncDatabase().

2. Kunci API Gemini:
Pastikan nilai GEMINI_API_KEY di dalam file .env sudah diisi dengan kunci API yang masih aktif. Tanpa kunci ini, fitur pencarian acara budaya pada rute /api/events dan /api/all tidak dapat menghasilkan data acara.
