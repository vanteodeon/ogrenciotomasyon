package student

import (
	"fmt"
	"ogrenci_otomasyon/database"
)

// Ekle, veritabanına yeni bir öğrenci kaydı oluşturur
func Ekle() {
	var no, ad, soyad, bolum string

	fmt.Println("\n--- Yeni Öğrenci Ekle ---")
	fmt.Print("Öğrenci No: ")
	fmt.Scan(&no)
	fmt.Print("Ad: ")
	fmt.Scan(&ad)
	fmt.Print("Soyad: ")
	fmt.Scan(&soyad)
	fmt.Print("Bölüm: ")
	fmt.Scan(&bolum)

	query := "INSERT INTO Ogrenciler (OgrenciNo, Ad, Soyad, Bolum) VALUES (@p1, @p2, @p3, @p4)"
	_, err := database.DB.Exec(query, no, ad, soyad, bolum)

	if err != nil {
		fmt.Println("[HATA] Kayıt işlemi başarısız oldu. Öğrenci numarası zaten mevcut olabilir.")
		return
	}
	fmt.Println("[BAŞARILI] Öğrenci sisteme eklendi!")
}

// Listele, veritabanındaki tüm öğrencileri ekrana yazdırır
func Listele() {
	fmt.Println("\n--- Kayıtlı Öğrenciler ---")
	query := "SELECT ID, OgrenciNo, Ad, Soyad, Bolum FROM Ogrenciler"
	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println("[HATA] Öğrenciler getirilirken bir sorun oluştu:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("%-5s | %-10s | %-15s | %-15s | %-20s\n", "ID", "Öğrenci No", "Ad", "Soyad", "Bölüm")
	fmt.Println("-----------------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var no, ad, soyad, bolum string
		err := rows.Scan(&id, &no, &ad, &soyad, &bolum)
		if err != nil {
			continue
		}
		fmt.Printf("%-5d | %-10s | %-15s | %-15s | %-20s\n", id, no, ad, soyad, bolum)
	}
}

// Sil, verilen öğrenci numarasına göre veritabanından kaydı siler
func Sil() {
	var no string
	fmt.Println("\n--- Öğrenci Sil ---")
	fmt.Print("Silinecek Öğrencinin Numarası: ")
	fmt.Scan(&no)

	query := "DELETE FROM Ogrenciler WHERE OgrenciNo = @p1"
	res, err := database.DB.Exec(query, no)

	if err != nil {
		fmt.Println("[HATA] Silme işlemi sırasında bir hata oluştu.")
		return
	}

	etkilenenKayit, _ := res.RowsAffected()
	if etkilenenKayit == 0 {
		fmt.Println("[UYARI] Bu numaraya sahip bir öğrenci bulunamadı.")
	} else {
		fmt.Println("[BAŞARILI] Öğrenci sistemden silindi!")
	}
}
