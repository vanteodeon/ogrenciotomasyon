package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

// ConnectDB veritabanı bağlantısını başlatır
func ConnectDB() {
	// SSMS bağlantı dizesi (Windows Authentication veya standart kurulumlar için)
	// Eğer SSMS'te farklı bir port, kullanıcı adı veya şifre kullanıyorsan burayı güncellemelisin.
	connString := "server=127.0.0.1;port=1433;database=OgrenciOtomasyonDB;trusted_connection=yes;"

	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Veritabanı bağlantı hatası: ", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Veritabanına ulaşılamıyor: ", err.Error())
	}

	fmt.Println("Veritabanı bağlantısı başarıyla kuruldu!")
}
