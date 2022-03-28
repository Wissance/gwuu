package gorm

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wissance/stringFormatter"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	g "gorm.io/gorm"
	"strings"
)

type SqlDialect string

const (
    Postgres SqlDialect = "postgres"
    Mysql = "mysql"
    Mssql = "mssql"
    Sqlite = "sqlite"
)

const postgresConnStrTemplate = "host={0} port={1} user={2} dbname={3} password={4} sslmode={5}"
const mssqlConnStrTemplate = "sqlserver://{username}:{password}@{host}:{port}?database={dbname}"
// todo: umv: think about charset as parameter
const mysqlConnStrTemplate = "{username}:{password}@tcp({host}:{port})/{dbname}?charset=utf8mb4&parseTime=True&loc=Local"
const postgresSystemDb = "postgres"
const mssqlSystemDb = "master"
const mysqlSystemDb = "mysql"

// BuildConnectionString
/* Function that builds connection string from individual parameters to use in OpenDb2
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - host - ip address / hostname of machine where database server is located
 *    - port - integer value representing server tcp port (typically 5432 for postgres, 3306 for mysql and 1433 for mssql)
 *    - dbName - database/catalog/schema name
 *    - dbUser - user that is using for perform operations on dbName
 *    - password - dbUser password
 *    - useSsl - is a string value that currently is using with Postgres Sql Only (allowed options are: disable, and others for enable)
 * Returns database connection string
 */
func BuildConnectionString(dialect SqlDialect, host string, port int, dbName string, dbUser string, password string, useSsl string) string {
	return createConnStr(dialect, host, port, dbName, dbUser, password, useSsl)
}

// OpenDb
/* Function that Open or Create and Open database
 * If you are using MSSQL Do not forget to switch on TCP connections for sql server, otherwise you wil get following error:
 * Unable to open tcp connection with host '127.0.0.1:1433': dial tcp 127.0.0.1:1433: connectex: No connection could
 * be made because the target machine actively refused it.
 * Ensure that You allowed port usage by Sql Server Connection Manager
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - host - ip address / hostname of machine where database server is located
 *    - port - integer value representing server tcp port (typically 5432 for postgres, 3306 for mysql and 1433 for mssql)
 *    - dbName - database/catalog/schema name
 *    - dbUser - user that is using for perform operations on dbName
 *    - password - dbUser password
 *    - create - if true we should create database if it does not exist
 *    - check - if true existence of database is checking otherwise not (we sure that this is a random database and to save
 *              some time we could omit existence check)
 *    - options - gorm config (from gorm.io/gorm not from github.com/jinzhu/gorm)
 * Returns gorm.DB address of database context object
 */
func OpenDb(dialect SqlDialect, host string, port int, dbName string, dbUser string, password string,
	        useSsl string, create bool, check bool, options *g.Config) *g.DB {
    connStr := createConnStr(dialect, host, port, dbName, dbUser, password, useSsl)
    return OpenDb2(dialect, connStr, create, check, options)
}

// OpenDb2
/* Function that Open or Create and Open database
 * This function does same as OpenDb but there is only one difference in parameters: for this function we pass connection string
 * instead of a lot of individual parameters
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - connStr - full connection string
 *    - create - if true we should create database if it does not exist
 *    - check - if true existence of database is checking otherwise not (we sure that this is a random database and to save
 *              some time we could omit existence check)
 *    - options - gorm config (from gorm.io/gorm not from github.com/jinzhu/gorm)
 */
func OpenDb2(dialect SqlDialect, connStr string, create bool, check bool, options *g.Config) *g.DB {
	dbCheckResult := CheckDb(dialect, connStr)
	if create == false {
		if dbCheckResult == false {
			return nil
		}
	} else {
		if !dbCheckResult {
			systemDbConnStr, dbName := createSystemDbConnStr(dialect, &connStr)
			return createDb(dialect, &systemDbConnStr, &connStr, &dbName, options)
		}
	}

	db, err := g.Open(createDialector(dialect, connStr), options)
	if err != nil{
		return nil
	}

	return db
}

// CheckDb
/* Functions that checks if database exists or not
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - connStr - full connection string
 * Returns true if database exists otherwise false
 */
func CheckDb(dialect SqlDialect, dbConnStr string) bool {
	db, err := g.Open(createDialector(dialect, dbConnStr), nil)
	if err == nil {
		sqlDb, err := db.DB()
		if err == nil && sqlDb != nil {
			err = sqlDb.Close()
		}
		return true
	}
	return false
}

// CloseDb
/* Function that close connection to database
 * Parameters:
 *    - db - address of database context object
 * Returns true on success
 */
func CloseDb(db *g.DB) bool {
	if db != nil {
		sqlDB, err := db.DB()
		if err == nil && sqlDB != nil {
			err = sqlDB.Close()
			if err == nil {
				return true
			}
		}
	}
	return false
}

// DropDb
/* Function that drops database from server
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - connStr - full connection string
 */
func DropDb(dialect SqlDialect, connStr string) bool {
	systemDbConnStr, dbName := createSystemDbConnStr(dialect, &connStr)
	return DropDb2(dialect, systemDbConnStr, dbName)
}

// DropDb2
/* Function that drop database from server using system database and dropping database name
 * Parameters:
 *     - dialect - string that represent using db driver inside gorm (see enum above)
 *     - systemDbConnStr - connection string to system database (in mysql - mysql, in sqlserver - master,
 *                         in postgres - postgres)
 *     - dbName - name of database that should be deleted
 * Returns true if database was deleted / dropped
 */
func DropDb2(dialect SqlDialect, systemDbConnStr string, dbName string) bool {
	db, err := g.Open(createDialector(dialect, systemDbConnStr), nil)
	if err != nil {
		return false
	}
	dropDbStatement := stringFormatter.Format("DROP DATABASE IF EXISTS {0}", dbName)
	err = db.Exec(dropDbStatement).Error
	if err != nil {
		CloseDb(db)
		return false
	}
	CloseDb(db)
    return true
}

// createSystemDbConnStr
/* Function that creates system database connection string from target database connection string
 * Create system db conn string using connection string to open target database, but database could not exist
 * therefore in some cases we have to create it (if we pass create=true to any OpenDb function).
 * In this function we are processing target db connStr and replace database name with system database name
 * Parameters:
 *     - dialect - string that represent using db driver inside gorm (see enum above)
 *     - connStr - connection string to other database
 * Return tuple of systemDbConnStr, dbName
 */
func createSystemDbConnStr(dialect SqlDialect, connStr *string) (string, string) {
	connStrCopy := *connStr
	if dialect == Postgres {
        // replace dbname={
		const postgresDbPattern = "dbname="
		beginIndex := strings.Index(connStrCopy, postgresDbPattern)
		if beginIndex < 0 {
			return "", ""
		}
		endIndex := getSymbolIndex(&connStrCopy, ' ', beginIndex +  len(postgresDbPattern))
		dbNameStr := connStrCopy[beginIndex: endIndex]
		systemDbStr := postgresDbPattern + postgresSystemDb
		return strings.Replace(connStrCopy, dbNameStr, systemDbStr, 1), dbNameStr[7:]

	} else if dialect == Mssql {
        const mssqlDbPattern = "?database="
		beginIndex := strings.Index(connStrCopy, mssqlDbPattern)
		if beginIndex < 0 {
			return "", ""
		}
		dbNameStr := connStrCopy[beginIndex:]
		systemDbStr := mssqlDbPattern + mssqlSystemDb
		return strings.Replace(connStrCopy, dbNameStr, systemDbStr, 1), dbNameStr[10:]

	} else if dialect == Mysql {
        beginIndex := getSymbolIndex(&connStrCopy, '/', 0)
        if beginIndex < 0 {
        	return "", ""
		}
		endIndex := getSymbolIndex(&connStrCopy, '?', beginIndex)
		dbNameStr := connStrCopy[beginIndex: endIndex]
		systemDbStr := "/" + mysqlSystemDb
		return strings.Replace(connStrCopy, dbNameStr, systemDbStr, 1), dbNameStr[1:]
	}
	return "", ""
}

// createConnStr
/* Function that creates connection string from individual parameters
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - host - ip address / hostname of machine where database server is located
 *    - port - integer value representing server tcp port (typically 5432 for postgres, 3306 for mysql and 1433 for mssql)
 *    - dbName - database/catalog/schema name
 *    - dbUser - user that is using for perform operations on dbName
 *    - password - dbUser password
 *    - useSsl - is a string value that currently is using with Postgres Sql Only (allowed options are: disable, and others for enable)
 * Returns connection string
 */
func createConnStr(dialect SqlDialect, host string, port int, dbName string,
	              dbUser string, password string, useSsl string) string {
	connStr := ""
	if dialect == Postgres {
        connStr = stringFormatter.Format(postgresConnStrTemplate, host, port, dbUser, dbName, password, useSsl)
	} else if dialect == Mssql {
        connStr = stringFormatter.FormatComplex(mssqlConnStrTemplate, map[string]interface{}{
        	"username":dbUser, "password":password, "host":host, "port":port, "dbname":dbName})
	} else if dialect == Mysql {
		connStr = stringFormatter.FormatComplex(mysqlConnStrTemplate, map[string]interface{}{
			"username":dbUser, "password":password, "host":host, "port":port, "dbname":dbName})
	}
	return connStr
}

// createDb
/* Function that creates database on server
 * Parameters:
 *    - dialect - string that represent using db driver inside gorm (see enum above)
 *    - systemDbConnStr - system (mysql - mysql, postgres - postgres, sqlserver - master) database connection string
 *    - dbConnStr - target database connection string
 *    - dbName - database name
 *    - options - gorm context configuration
 * Return pointer to database context
 */
func createDb(dialect SqlDialect, systemDbConnStr *string, dbConnStr *string, dbName *string, options *g.Config) *g.DB {
	createStatementTemplate := "CREATE DATABASE {0}"
	createStatement := stringFormatter.Format(createStatementTemplate, *dbName)

	systemDb, err := g.Open(createDialector(dialect, *systemDbConnStr), nil)
	if err != nil {
		return nil
	}
	systemDb.Exec(createStatement)
	sqlDb, err := systemDb.DB()
	if err == nil {
		sqlDb.Close()
	}
	db, err := g.Open(createDialector(dialect, *dbConnStr), options)
	if err != nil {
		return nil
	}
	return db
}

// getSymbolIndex
/* Function that searches index of symbol in string from start position (index)
 * Parameters:
 *    - str - string where we are searching for a symbol
 *    - symbol - single symbol that we are searching in a string
 *    - startIndex - index of string from what we
 * Returns index of symbol in str otherwise -1
 */
func getSymbolIndex(str *string, symbol rune, startIndex int) int {
	strSymbols := []rune(*str)
	if startIndex < 0 {
		startIndex = 0
	}
	for i := startIndex; i < len(*str); i++ {
		if strSymbols[i] == symbol {
            return i
		}
	}
	return  -1
}

// createDialector
/* Function that creates dialector (calls Open of driver)
 * Parameters:
 *    - dialect - dialect of database server
 *    - dbConnStr -
 * Return dialector or nil
 */
func createDialector(dialect SqlDialect, dbConnStr string) g.Dialector {
	if dialect == Mysql {
		return mysql.Open(dbConnStr)
	}
	if dialect == Mssql {
        return sqlserver.Open(dbConnStr)
	}
	if dialect == Postgres {
		return postgres.Open(dbConnStr)
	}
    //return sqlite.Open(dbConnStr)
    return nil
}
