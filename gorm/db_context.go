package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/wissance/stringFormatter"
	"strings"
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
const postgresSystemDb = "postgres"
const mssqlSystemDb = "master"
const mysqlSystemDb = "mysql"


func OpenDb(dialect SqlDialect, host string, port int, dbName string, dbUser string, password string,
	        useSsl string, create bool) *gorm.DB {
    connStr := createConnStr(dialect, host, port, dbName, dbUser, password, useSsl)
    return OpenDb2(dialect, connStr, create)
}

func OpenDb2(dialect SqlDialect, connStr string, create bool) *gorm.DB {
	dbCheckResult := CheckDb(dialect, connStr)
	if create == false {
		if dbCheckResult == false {
			return nil
		}
	} else {

	}
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

/*
 * Create system db conn string using connection string to open target database, but database could not exists
 * therefore in some cases we have to create it (if we pass create=true to any OpenDb function).
 * In this function we are processing target db connStr and replace database name with system database name
 */
func createSystemDbConnStr(dialect SqlDialect, connStr *string) string {
	connStrCopy := *connStr
	if dialect == Postgres {
        // replace dbname={
		const postgresDbPattern = "dbname="
		beginIndex := strings.Index(connStrCopy, postgresDbPattern)
		if beginIndex < 0 {
			return ""
		}
		endIndex := getSymbolIndex(&connStrCopy, ' ', beginIndex +  len(postgresDbPattern))
		dbNameStr := connStrCopy[beginIndex: endIndex]
		systemDbStr := postgresDbPattern + postgresSystemDb
		return strings.Replace(connStrCopy, dbNameStr, systemDbStr, 1)

	} else if dialect == Mssql {

	} else if dialect == Mysql {

	}
	return ""
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

func getSymbolIndex(str *string, symbol rune, startIndex int) int {
	strSymbols := []rune(*str)
	for i := startIndex; i < len(*str); i++ {
		if strSymbols[i] == symbol {
            return i;
		}
	}
	return  -1
}