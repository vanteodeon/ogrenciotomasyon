package auth

import (
	"database/sql"
	"fmt"
	"ogrenci_otomasyon/database" // Kendi modül adımıza göre database paketini çağırıyoruz
)

// Login fonksiyonu, kullanıcıdan bilgi alıp veritabanında doğrular
func Login() bool {
	var username, password string

	fmt.Println("=====================================")
	fmt.Println("  ÖĞRENCİ OTOMASYON SİSTEMİ GİRİŞİ   ")
	fmt.Println("=====================================")

	fmt.Print("Kullanıcı Adı: ")
	fmt.Scanln(&username)

	fmt.Print("Şifre: ")
	fmt.Scanln(&password)

	var dbUsername string
	// Veritabanında bu kullanıcı adı ve şifreyle eşleşen bir kayıt arıyoruz
	query := "SELECT KullaniciAdi FROM Kullanicilar WHERE KullaniciAdi = @p1 AND Sifre = @p2"
	err := database.DB.QueryRow(query, username, password).Scan(&dbUsername)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("\n[HATA] Kullanıcı adı veya şifre yanlış!\n")
			return false
		}
		fmt.Println("\n[HATA] Veritabanı sorgu hatası:", err)
		return false
	}

	fmt.Printf("\n[BAŞARILI] Giriş onaylandı! Hoş geldiniz, %s.\n\n", dbUsername)
	return true
}
