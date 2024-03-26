package config_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"sekolahbeta/miniproject3/config"
	"testing"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

func TestKoneksi(t *testing.T) {
	Init()
	config.OpenDB()
}
