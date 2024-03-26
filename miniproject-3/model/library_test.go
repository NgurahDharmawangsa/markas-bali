package model_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"sekolahbeta/miniproject3/config"
	"sekolahbeta/miniproject3/model"
	"testing"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}

func TestCreateCarSuccess(t *testing.T) {
	Init()

	libraryData := model.Library{
		ISBN:    "123",
		Penulis: "Ahmad S",
		Tahun:   2020,
		Judul:   "Selamat",
		Gambar:  "gambar1.jpg",
		Stok:    2,
	}

	err := libraryData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	//fmt.Println(libraryData)

	//config.Mysql.DB.Unscoped().Delete(&libraryData)
}

func TestGetByIdSuccess(t *testing.T) {
	Init()

	libraryData := model.Library{
		ISBN:    "6767545",
		Penulis: "Hendrawan Teja Bukti",
		Tahun:   2006,
		Judul:   "Manusia Setengah Dua Belas",
		Gambar:  "12.jpg",
		Stok:    1,
	}

	err := libraryData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	libraryData = model.Library{
		Model: model.Model{
			ID: libraryData.ID,
		},
	}

	_, err = libraryData.GetById(config.Mysql.DB, libraryData.ID)
	assert.Nil(t, err)

	//fmt.Println(data)
	//config.Mysql.DB.Unscoped().Delete(&libraryData)
}

func TestGetAll(t *testing.T) {
	Init()

	libraryData := model.Library{
		ISBN:    "123",
		Penulis: "Sucipto Tejo",
		Tahun:   2018,
		Judul:   "Akulah Sang",
		Gambar:  "sku.jpg",
		Stok:    20,
	}

	err := libraryData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := libraryData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)

	//config.Mysql.DB.Unscoped().Delete(&libraryData)
}

func TestUpdateByID(t *testing.T) {
	Init()

	libraryData := model.Library{
		ISBN:    "222A9",
		Penulis: "Wahyu Kurniawan",
		Tahun:   2010,
		Judul:   "Sapi Gembala",
		Gambar:  "sapi.png",
		Stok:    22,
	}

	err := libraryData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	libraryData.Judul = "Sekar Jagat"

	err = libraryData.UpdateOneByID(config.Mysql.DB, libraryData.ID)
	assert.Nil(t, err)

	//config.Mysql.DB.Unscoped().Delete(&libraryData)
}

func TestDeleteByID(t *testing.T) {
	Init()

	libraryData := model.Library{
		ISBN:    "QWE12",
		Penulis: "Jajang Suherman",
		Tahun:   2019,
		Judul:   "Spiderman",
		Gambar:  "man.png",
		Stok:    120,
	}

	err := libraryData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	libraryData = model.Library{
		Model: model.Model{
			ID: libraryData.ID,
		},
	}

	err = libraryData.DeleteByID(config.Mysql.DB, libraryData.ID)
	assert.Nil(t, err)

	//config.Mysql.DB.Unscoped().Delete(&libraryData)
}
