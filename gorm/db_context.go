package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/wissance/stringFormatter"
)

type SqlDialect string

const (
    Postgres SqlDialect = "postgres"
    Mysql = "mysql"
    Mssql = "mssql"
    Sqlite = "sqlite"
)

const postrgresConnStrTemplate = "host={0} port={1} user={2} dbname={3} password={4} sslmode={5}"
const mssqlConnStrTemplate = "sqlserver://{username}:{password}@{host}:{port}?database={dbname}"
// todo: umv: think about charset as parameter
const mysqlConnStrTemplate = "{username}:{password}@tcp({host}:{port})/{dbname}?charset=utf8mb4&parseTime=True&loc=Local"

func OpenDb(dialect SqlDialect, host string, port string, dbName string, dbUser string, password string,
	        useSsl string, create bool) *gorm.DB {
    return nil
}

func OpenDb2(dialect SqlDialect, connStr string, create bool) *gorm.DB {
	return nil
}

func CheckDb(dialect SqlDialect, dbConnStr string) bool {
	db, err := gorm.Open(string(dialect), dbConnStr)
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

func createConnStr(dialect SqlDialect, host string, port int, dbName string,
	              dbUser string, password string, useSsl string) string {
	connStr := ""
	if dialect == Postgres {
        connStr = stringFormatter.Format(postrgresConnStrTemplate, host, port, dbUser, dbName, password, useSsl)
	} else if dialect == Mssql {
        connStr = stringFormatter.FormatComplex(mssqlConnStrTemplate, map[string]interface{}{
        	"username":dbUser, "password":password, "host":host, "port":port, "dbname":dbName})
	} else if dialect == Mysql {
		connStr = stringFormatter.FormatComplex(mysqlConnStrTemplate, map[string]interface{}{
			"username":dbUser, "password":password, "host":host, "port":port, "dbname":dbName})
	} else if dialect == Sqlite {

	}

	return connStr
}

func createDb(dialect SqlDialect, systemDbConnStr *string, dbConnStr *string, dbName *string) *gorm.DB {
	createStatementTemplate := "CREATE DATABASE {0}"
	createStatement := stringFormatter.Format(createStatementTemplate, *dbName)

	systemDb, err := gorm.Open(string(dialect), *systemDbConnStr)
	if err != nil {
		return nil
	}
	systemDb.Exec(createStatement)
	systemDb.Close()
	db, err := gorm.Open(string(dialect), *dbConnStr)
	if err != nil {
		return nil
	}
	return db
}