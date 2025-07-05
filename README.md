# api manajemen kinik

API Backend untuk manajemen data pasien, dokter, dan jadwal praktik. Dibuat dengan Golang dan SQLite. Siap untuk digunakan di sistem rumah sakit/klinik digital.

## 📦 Fitur

- CRUD Pasien
- CRUD Dokter
- Lihat Jadwal Dokter
- Struktur folder profesional
- Koneksi database otomatis

## 🚀 Instalasi

```bash
git clone https://github.com/yourusername/hospital-api-go-pro.git
cd hospital-api-go-pro
go mod tidy
go run main.go
```

## 🧪 Contoh Endpoint

| Method | Endpoint             | Keterangan              |
|--------|----------------------|--------------------------|
| GET    | /                    | Cek koneksi API         |
| GET    | /patients            | Lihat semua pasien      |
| POST   | /patients            | Tambah pasien           |
| GET    | /doctors             | Lihat semua dokter      |
| POST   | /doctors             | Tambah dokter baru      |
| GET    | /doctors/schedule    | Lihat jadwal dokter     |

## 📂 Struktur Folder

```
hospital-api-go-pro/
├── controllers/
├── models/
├── routes/
├── database/
├── middleware/ (belum dipakai)
├── main.go
├── go.mod
```

## 🧠 Rencana Fitur Lanjutan

- Login JWT
- Sistem antrian poli
- Input resep dan tindakan
- Billing dan invoice
- Export PDF
- Integrasi notifikasi Telegram

## 🪪 Lisensi

MIT License.
