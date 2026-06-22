package main

import (
	"fmt"
	"ogrenci_otomasyon/auth"
	"ogrenci_otomasyon/database"
	"ogrenci_otomasyon/student"
	"os"
)

func main() {
	// 1. Veritabanına Bağlan
	database.ConnectDB()

	// 2. Kimlik Doğrulama (Login) Ekranı
	if !auth.Login() {
		fmt.Println("Sisteme giriş yapılamadı. Program kapatılıyor.")
		os.Exit(1)
	}

	// 3. Ana Menü Döngüsü
	for {
		fmt.Println("\n=====================================")
		fmt.Println("             ANA MENÜ                ")
		fmt.Println("=====================================")
		fmt.Println("1. Yeni Öğrenci Ekle")
		fmt.Println("2. Öğrencileri Listele")
		fmt.Println("3. Öğrenci Sil")
		fmt.Println("0. Çıkış Yap")
		fmt.Print("Seçiminiz: ")

		var secim string
		fmt.Scan(&secim)

		switch secim {
		case "1":
			student.Ekle()
		case "2":
			student.Listele()
		case "3":
			student.Sil()
		case "0":
			fmt.Println("Sistemden çıkış yapılıyor... İyi günler!")
			os.Exit(0)
		default:
			fmt.Println("[HATA] Geçersiz seçim! Lütfen menüdeki numaralardan birini giriniz.")
		}
	}
}
