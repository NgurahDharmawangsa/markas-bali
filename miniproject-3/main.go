package main

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"sekolahbeta/miniproject3/config"
	"sekolahbeta/miniproject3/model"
	"sekolahbeta/miniproject3/utils"
	"sync"
	"time"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

func main() {
	Init()
	config.OpenDB()

	fmt.Println("\n")
	var opsi int

	fmt.Println("======")
	fmt.Println("Manajemen Buku Perpustakaan")
	fmt.Println("======")

	fmt.Println("Pilih Opsi")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Semua Buku")
	fmt.Println("3. Lihat Detail Buku")
	fmt.Println("4. Edit Buku")
	fmt.Println("5. Hapus Buku")
	fmt.Println("6. Generate PDF")
	fmt.Println("7. Import CSV ke Database")
	fmt.Println("8. Keluar")

	fmt.Print("Masukkan Opsi : ")
	_, err := fmt.Scanln(&opsi)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	switch opsi {
	case 1:
		tambahBuku(config.Mysql.DB)
	case 2:
		listBuku(config.Mysql.DB)
	case 3:
		var pilihDetail uint
		listBuku(config.Mysql.DB)
		fmt.Print("Masukkan ID Buku : ")
		_, err := fmt.Scanln(&pilihDetail)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}
		detailBuku(config.Mysql.DB, pilihDetail)
	case 4:
		var pilihanUpdate uint
		listBuku(config.Mysql.DB)
		fmt.Print("Masukkan ID Buku Yang Akan DiEdit : ")
		_, err := fmt.Scanln(&pilihanUpdate)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}
		updateBuku(config.Mysql.DB, pilihanUpdate)
	case 5:
		var pilihanHapus uint
		listBuku(config.Mysql.DB)
		fmt.Print("Masukkan Kode Buku Yang Akan Dihapus : ")
		_, err := fmt.Scanln(&pilihanHapus)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}
		hapusBuku(config.Mysql.DB, pilihanHapus)
	case 6:
		GeneratePdfBuku(config.Mysql.DB)
	case 7:
		err := utils.GetLibrary(config.Mysql.DB)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}
	case 8:
		os.Exit(0)
	default:
		fmt.Println("Tidak Ada Opsi")
	}

	main()
}

func tambahBuku(db *gorm.DB) {
	fmt.Println("\n")
	isbn := ""
	bookAuthor := ""
	var publishedYear uint
	bookTitle := ""
	bookImage := ""
	var stock uint

	fmt.Println("======")
	fmt.Println("Tambah Buku")
	fmt.Println("======")

	draftBuku := []model.Library{}

	for {
		fmt.Print("Masukkan ISBN : ")
		_, err := fmt.Scanln(&isbn)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}

		fmt.Print("Masukkan Penulis Buku : ")
		_, err = fmt.Scanln(&bookAuthor)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}

		fmt.Print("Masukkan Tahun Terbit : ")
		_, err = fmt.Scanln(&publishedYear)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}

		fmt.Print("Masukkan Judul Buku : ")
		_, err = fmt.Scanln(&bookTitle)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}

		fmt.Print("Masukkan Gambar Buku : ")
		_, err = fmt.Scanln(&bookImage)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}

		fmt.Print("Masukkan Stok Buku : ")
		_, err = fmt.Scanln(&stock)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}

		draftBuku = append(draftBuku, model.Library{
			ISBN:    isbn,
			Penulis: bookAuthor,
			Tahun:   publishedYear,
			Judul:   bookTitle,
			Gambar:  bookImage,
			Stok:    stock,
		})

		pilihanTambahBuku := 0
		fmt.Println("Ketik 1 untuk tambah buku, ketik 0 untuk keluar")
		_, err = fmt.Scanln(&pilihanTambahBuku)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		if pilihanTambahBuku == 0 {
			break
		}
	}

	fmt.Println("Menambah Buku...")

	_ = os.Mkdir("buku", 0777)

	ch := make(chan model.Library)

	wg := sync.WaitGroup{}

	jumlahPustakawan := 5

	for i := 0; i < jumlahPustakawan; i++ {
		wg.Add(1)
		go simpanBuku(ch, &wg, i, db)
	}

	for _, bukuTersimpan := range draftBuku {
		ch <- bukuTersimpan
	}

	close(ch)

	wg.Wait()

	fmt.Println("Berhasil Tambah Buku")
}

func simpanBuku(ch <-chan model.Library, wg *sync.WaitGroup, noPustakawan int, db *gorm.DB) {

	for bukuTersimpan := range ch {
		if err := db.Create(&bukuTersimpan).Error; err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		fmt.Printf("Pustakawan No %d Memproses Kode Buku : %s!\n", noPustakawan, bukuTersimpan.Judul)
	}
	wg.Done()
}

func lihatListBuku(ch <-chan string, chBuku chan model.Library, wg *sync.WaitGroup, db *gorm.DB) {
	var bookLibrary model.Library
	for bookISBN := range ch {
		if err := db.Where("isbn = ?", bookISBN).Find(&bookLibrary).Error; err != nil {
			fmt.Println("Terjadi error res:", err)
			continue
		}

		chBuku <- bookLibrary
	}
	wg.Done()
}

func listBuku(db *gorm.DB) {
	//var libraries []model.Library
	libraryData := model.Library{}
	res, err := libraryData.GetAll(db)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	fmt.Println("\n")
	fmt.Println("======")
	fmt.Println("List Buku")
	fmt.Println("======")
	fmt.Println("Memuat Data...")

	model.ListBook = res

	wg := sync.WaitGroup{}
	ch := make(chan string)
	chBuku := make(chan model.Library, len(model.ListBook))

	jumlahPustakawan := 5

	for i := 0; i < jumlahPustakawan; i++ {
		wg.Add(1)
		go lihatListBuku(ch, chBuku, &wg, db)
	}

	for _, library := range res {
		ch <- library.ISBN
	}
	close(ch)

	wg.Wait()
	close(chBuku)

	//for dataBookLibrary := range chBuku {
	//	model.ListBook = append(model.ListBook, dataBookLibrary)
	//}

	if len(model.ListBook) < 1 {
		fmt.Println("---Tidak Ada Buku---")
	}

	for i, v := range model.ListBook {
		i++
		fmt.Printf("%d. ID : %d, ISBN : %s, Penulis : %s, Tahun : %d, Judul : %s, Gambar : %s, Stok : %d\n", i, v.ID, v.ISBN, v.Penulis, v.Tahun, v.Judul, v.Gambar, v.Stok)
	}
}

func detailBuku(db *gorm.DB, id uint) {
	fmt.Println("\n")
	fmt.Println("======")
	fmt.Println("Detail Buku")
	fmt.Println("======")

	libraryData := model.Library{}

	book, err := libraryData.GetById(db, id)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	fmt.Printf("ISBN : %s\n", book.ISBN)
	fmt.Printf("Penulis : %s\n", book.Penulis)
	fmt.Printf("Tahun : %d\n", book.Tahun)
	fmt.Printf("Judul : %s\n", book.Judul)
	fmt.Printf("Gambar : %s\n", book.Gambar)
	fmt.Printf("Stok : %d\n", book.Stok)
}

func updateBuku(db *gorm.DB, id uint) {
	fmt.Println("\n")
	detailBuku(db, id)

	fmt.Println("======")
	fmt.Println("Edit Buku")
	fmt.Println("======")

	var book model.Library

	err := db.Where("id = ?", id).First(&book).Error
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return
	}

	fmt.Print("Masukkan ISBN : ")
	_, err = fmt.Scanln(&book.ISBN)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Penulis : ")
	_, err = fmt.Scanln(&book.Penulis)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Tahun : ")
	_, err = fmt.Scanln(&book.Tahun)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Judul : ")
	_, err = fmt.Scanln(&book.Judul)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Gambar : ")
	_, err = fmt.Scanln(&book.Gambar)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Stok : ")
	_, err = fmt.Scanln(&book.Stok)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	err = book.UpdateOneByID(db, book.ID)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat melakukan update:", err)
		return
	}

	fmt.Println("Buku berhasil diupdate")

}

func hapusBuku(db *gorm.DB, id uint) {
	fmt.Println("\n")
	var isBook bool

	for _, book := range model.ListBook {
		if book.ID == id {
			isBook = true

			err := book.DeleteByID(db, id)
			if err != nil {
				fmt.Println("Terjadi error:", err)
				return
			}

			fmt.Println("Buku Berhasil Dihapus")
			break
		}
	}

	if !isBook {
		fmt.Println("Kode Buku Salah Atau Tidak Ada")
	}
}

func GeneratePdfBuku(db *gorm.DB) {
	listBuku(db)
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Buku ...")
	fmt.Println("=================================")

	//fmt.Println(model.ListBook)
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for i, buku := range model.ListBook {
		bukuText := fmt.Sprintf(
			"Buku #%d:\nID : %d\nISBN : %s\nPenulis : %s\nTahun : %d\nJudul : %s\nGambar : %s\nStok : %d\n",
			i+1, buku.ID, buku.ISBN,
			buku.Penulis, buku.Tahun, buku.Judul, buku.Gambar,
			buku.Stok)

		pdf.MultiCell(0, 10, bukuText, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}
