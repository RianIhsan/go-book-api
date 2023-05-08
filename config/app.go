package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {

	d, err := gorm.Open("mysql", "root:@/bookrestapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Gagal koneksi ke database")
	} else {
		fmt.Println("Berhasil terkoneksi")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
