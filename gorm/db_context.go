package gorm

import (
	/*
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wissance/stringFormatter"*/
	"github.com/jinzhu/gorm"
)

func OpenDb(host string, port string, dbName string, dbUser string, password string,
	        useSsl string, create bool) *gorm.DB {
    return nil
}

func OpenDb2(connStr string, create bool) *gorm.DB {
	return nil
}

func CheckDb(dbConnStr *string) bool {
	db, err := gorm.Open("postgres", *dbConnStr)
	if err == nil {
		db.Close()
		return true
	}
	return false
}

func CloseDb(db *gorm.DB) {
	if db != nil {
		defer db.Close()
	}
}

func DropDb(systemDbConnStr string, dbName string, checkExists bool) {

}