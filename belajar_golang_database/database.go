package belajargolangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	// Setting Connection Pooling Database
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)  // Suatu ketika sistem tersebut berhenti bekerja atau bengong seleama 5 minute, maka secara otomatis sistem akan mengclose dari koneksi aplikasi ke database
	db.SetConnMaxLifetime(60 * time.Minute) // Suatu ketika sistem tersebut sedang mengkoneksikan ke database dan terjadi interval waktu mencapai 60 menit dan belum terdapat respon balik dari database, maka akan dibuatkan kembali koneksi yang baru

	return db
}
