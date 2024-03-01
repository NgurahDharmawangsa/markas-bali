package main

import (
	"fmt"
	"os"
)

type BookLibrary struct {
	KodeBuku      string
	JudulBuku     string
	Pengarang     string
	Penerbit      string
	JumlahHalaman int
	TahunTerbit   int
}

var ListBook []BookLibrary

func tambahBuku() {
	fmt.Println("\n")
	bookCode := ""
	bookTitle := ""
	bookAuthor := ""
	bookPublisher := ""
	var pageTotal int
	var publishedYear int

	fmt.Println("======")
	fmt.Println("Tambah Buku")
	fmt.Println("======")

	fmt.Print("Masukkan Code Buku : ")
	_, err := fmt.Scanln(&bookCode)
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

	fmt.Print("Masukkan Pengarang Buku : ")
	_, err = fmt.Scanln(&bookAuthor)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Penerbit Buku : ")
	_, err = fmt.Scanln(&bookPublisher)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Total Halaman : ")
	_, err = fmt.Scanln(&pageTotal)
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

	ListBook = append(ListBook, BookLibrary{
		KodeBuku:      bookCode,
		JudulBuku:     bookTitle,
		Pengarang:     bookAuthor,
		Penerbit:      bookPublisher,
		JumlahHalaman: pageTotal,
		TahunTerbit:   publishedYear,
	})

	fmt.Println("Berhasil Tambah Buku")
}

func listBuku() {
	fmt.Println("\n")
	fmt.Println("======")
	fmt.Println("List Buku")
	fmt.Println("======")

	if len(ListBook) < 1 {
		fmt.Println("---Tidak Ada Buku---")
	}

	for i, v := range ListBook {
		i++
		fmt.Printf("%d. Kode Buku : %s, Judul Buku : %s, Pengarang : %s, Penerbit : %s, Jumlah Halaman : %d, Tahun Terbit : %d\n", i, v.KodeBuku, v.JudulBuku, v.Pengarang, v.Penerbit, v.JumlahHalaman, v.TahunTerbit)
	}
}

func detailBuku(kode string) {
	fmt.Println("\n")
	fmt.Println("======")
	fmt.Println("Detail Buku")
	fmt.Println("======")

	var isBook bool

	for _, book := range ListBook {
		if book.KodeBuku == kode {
			isBook = true
			fmt.Printf("Kode Buku : %s\n", book.KodeBuku)
			fmt.Printf("Judul Buku : %s\n", book.JudulBuku)
			fmt.Printf("Pengarang Buku : %s\n", book.Pengarang)
			fmt.Printf("Penerbit Buku : %s\n", book.Penerbit)
			fmt.Printf("Jumlah Halaman : %d\n", book.JumlahHalaman)
			fmt.Printf("Tahun Terbit : %d\n", book.TahunTerbit)
			break
		}
	}

	if !isBook {
		fmt.Println("Kode Buku Salah Atau Tidak Ada")
	}
}

func updateBuku(kode string) {
	fmt.Println("\n")
	detailBuku(kode)

	fmt.Println("======")
	fmt.Println("Edit Buku")
	fmt.Println("======")

	var book BookLibrary

	fmt.Print("Masukkan Code Buku : ")
	_, err := fmt.Scanln(&book.KodeBuku)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Judul Buku : ")
	_, err = fmt.Scanln(&book.JudulBuku)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Pengarang Buku : ")
	_, err = fmt.Scanln(&book.Pengarang)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Penerbit Buku : ")
	_, err = fmt.Scanln(&book.Penerbit)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Total Halaman : ")
	_, err = fmt.Scanln(&book.JumlahHalaman)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Print("Masukkan Tahun Terbit : ")
	_, err = fmt.Scanln(&book.TahunTerbit)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	fmt.Println(book)

	for i, b := range ListBook {
		if b.KodeBuku == kode {
			ListBook[i] = book
			break
		}
	}
}

func hapusBuku(kode string) {
	fmt.Println("\n")
	for i, book := range ListBook {
		if book.KodeBuku == kode {
			ListBook = append(ListBook[:i], ListBook[i+1:]...)
			fmt.Println("Buku Berhasil Dihapus")
			break
		}
	}

}

func main() {
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
	fmt.Println("6. Keluar")

	fmt.Print("Masukkan Opsi : ")
	_, err := fmt.Scanln(&opsi)
	if err != nil {
		fmt.Println("Terjadi Kesalahan : ", err)
		return
	}

	switch opsi {
	case 1:
		tambahBuku()
	case 2:
		listBuku()
	case 3:
		var pilihDetail string
		listBuku()
		fmt.Print("Masukkan Kode Buku : ")
		_, err := fmt.Scanln(&pilihDetail)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}
		detailBuku(pilihDetail)
	case 4:
		var pilihanUpdate string
		listBuku()
		fmt.Print("Masukkan Kode Buku Yang Akan DiEdit : ")
		_, err := fmt.Scanln(&pilihanUpdate)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}
		updateBuku(pilihanUpdate)
	case 5:
		var pilihanHapus string
		listBuku()
		fmt.Print("Masukkan Kode Buku Yang Akan Dihapus : ")
		_, err := fmt.Scanln(&pilihanHapus)
		if err != nil {
			fmt.Println("Terjadi Kesalahan : ", err)
			return
		}
		hapusBuku(pilihanHapus)
	case 6:
		os.Exit(0)
	default:
		fmt.Println("Tidak Ada Opsi")
	}

	main()
}
