package utils

import (
	"encoding/csv"
	"fmt"
	"gorm.io/gorm"
	"os"
	"sekolahbeta/miniproject3/model"
	"strconv"
	"sync"
	"time"
)

func openFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetLibrary(db *gorm.DB) error {
	started := time.Now()
	file, err := os.Open("sample_books.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer file.Close()

	csvChan, err := loadFileGoroutine(file)
	if err != nil {
		return err
	}

	jmlGoroutine := 10

	var libChanTemp []<-chan model.Library

	for i := 0; i < jmlGoroutine; i++ {
		libChanTemp = append(libChanTemp, processConvertStruct(csvChan, db))
	}

	appendLibrary(libChanTemp...)

	fmt.Println("[Dengan Goroutine]", time.Since(started))

	return nil
}

func loadFileGoroutine(file *os.File) (<-chan []string, error) {
	libraryChan := make(chan []string)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return libraryChan, err
	}

	go func() {
		for _, library := range records {
			libraryChan <- library
		}

		close(libraryChan)
	}()

	return libraryChan, nil

}

func processConvertStruct(csvChan <-chan []string, db *gorm.DB) <-chan model.Library {
	librariesChan := make(chan model.Library)

	go func() {
		for libraryData := range csvChan {
			idu, err := strconv.Atoi(libraryData[0])
			if err != nil {
				fmt.Println("Terjadi kesalahan konversi ID:", err)
				continue
			}
			tahunu, err := strconv.Atoi(libraryData[3])
			if err != nil {
				fmt.Println("Terjadi kesalahan konversi Tahun:", err)
				continue
			}
			stoku, err := strconv.Atoi(libraryData[6])
			if err != nil {
				fmt.Println("Terjadi kesalahan konversi Stok:", err)
				continue
			}

			library := model.Library{
				Model: model.Model{
					ID:        uint(idu),
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				ISBN:    libraryData[1],
				Penulis: libraryData[2],
				Tahun:   uint(tahunu),
				Judul:   libraryData[4],
				Gambar:  libraryData[5],
				Stok:    uint(stoku),
			}

			if err := db.Create(&library).Error; err != nil {
				fmt.Println("Failed to save library to database:", err)
				continue
			}

			librariesChan <- library
		}

		close(librariesChan)
	}()

	return librariesChan
}

func appendLibrary(libraryChanMany ...<-chan model.Library) {
	wg := sync.WaitGroup{}

	wg.Add(len(libraryChanMany))
	for _, ch := range libraryChanMany {
		go func(ch <-chan model.Library) {
			for range ch {

			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
	}()

	return
}
