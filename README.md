# Books API

Books API adalah RESTful API sederhana yang dibuat menggunakan Golang dan Gin Framework. Project ini dibuat sebagai latihan untuk memahami dasar pengembangan backend, mulai dari routing, controller, model, koneksi database, hingga operasi CRUD menggunakan GORM.

## Features

- Menambahkan data buku
- Menampilkan seluruh data buku
- Menampilkan detail buku berdasarkan ID
- Mengubah data buku
- Menghapus data buku

## Tech Stack

- Golang
- Gin Framework
- GORM
- MySQL
- godotenv

## Project Structure

```
books-api
├── config
├── controllers
├── models
├── routers
├── main.go
├── .env
└── go.mod
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/book` | Menambahkan buku baru |
| GET | `/book` | Menampilkan semua buku |
| GET | `/book/:id` | Menampilkan detail buku |
| PUT | `/book/:id` | Mengubah data buku |
| DELETE | `/book/:id` | Menghapus data buku |

## Installation

Clone repository

```bash
git clone https://github.com/RafiBakhtiar23/books-api.git
```

Masuk ke folder project

```bash
cd books-api
```

Install dependency

```bash
go mod tidy
```

Buat file `.env`

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=books_db
```

Jalankan project

```bash
go run main.go
```

Server berjalan di

```
http://localhost:8080
```

## Database

Project ini menggunakan MySQL dengan GORM sebagai ORM. Tabel akan dibuat otomatis saat aplikasi dijalankan melalui fitur AutoMigrate.

## Notes

Project ini dibuat untuk latihan dasar pengembangan REST API menggunakan Golang dan sebagai bagian dari portfolio backend.
