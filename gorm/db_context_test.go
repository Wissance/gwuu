package gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbUser = "developer"
const dbPassword = "123"

// ########################################### public functions tests #################################################

// test Build connection string

func TestBuildPostgresConnectionString(t *testing.T) {
    connStr := BuildConnectionString(Postgres, "localhost", 5432, "custom_app",
    	                             "root", "P@ssW0rd", "disable")
    expectedConnStr := "host=localhost port=5432 user=root dbname=custom_app password=P@ssW0rd sslmode=disable"
    assert.Equal(t, expectedConnStr, connStr)
}

func TestBuildMysqlConnectionString(t *testing.T) {
	connStr := BuildConnectionString(Mysql, "127.0.0.1", 3306, "custom_app",
		                             "root", "P@ssW0rd", "")
	expectedConnStr := "root:P@ssW0rd@tcp(127.0.0.1:3306)/custom_app?charset=utf8mb4&parseTime=True&loc=Local"
	assert.Equal(t, expectedConnStr, connStr)
}

func TestBuildMssqlConnectionString(t *testing.T) {
	connStr := BuildConnectionString(Mssql, "192.168.10.100", 1433, "custom_app",
		                            "sa", "123", "")
	expectedConnStr := "sqlserver://sa:123@192.168.10.100:1433?database=custom_app"
	assert.Equal(t, expectedConnStr, connStr)
}

// test open db (system db without create)

func TestPostgresOpenSystemDb(t *testing.T) {
    db := OpenDb(Postgres, "127.0.0.1", 5432, "postgres", dbUser, dbPassword, "disable", false)
    assert.NotNil(t, db)
    CloseDb(db)
}

func TestMysqlOpenSystemDb(t *testing.T) {
	db := OpenDb(Mysql, "localhost", 3306, "mysql", dbUser, dbPassword, "", false)
	assert.NotNil(t, db)
	CloseDb(db)
}

func TestMssqlOpenSystemDb(t *testing.T) {
	db := OpenDb(Mssql, "localhost", 1433, "master", dbUser, dbPassword, "", false)
	assert.NotNil(t, db)
	CloseDb(db)
}

// test open db with create
func TestPostgresOpenDbWithCreate(t *testing.T) {
    // Create Db when open
	connStr := BuildConnectionString(Postgres, "127.0.0.1", 5432, "gwuu_examples", dbUser, dbPassword, "disable")
	testOpenDbWithCreateAndCheck(t, connStr, Postgres)
}

func TestMysqlOpenDbWithCreate(t *testing.T) {
	connStr := BuildConnectionString(Mysql, "127.0.0.1", 3306, "gwuu_examples", dbUser, dbPassword, "")
	testOpenDbWithCreateAndCheck(t, connStr, Mysql)
}

func TestMssqlOpenDbWithCreate(t *testing.T) {
	connStr := BuildConnectionString(Mssql, "localhost", 1433, "GwuuExamples", dbUser, dbPassword, "")
	testOpenDbWithCreateAndCheck(t, connStr, Mssql)
}

// ####################################################################################################################

// ########################################### private functions tests ################################################

func TestCreatePostgresSystemDbConnectionString(t *testing.T) {
    connStr := "host=localhost port=5432 user=root dbname=custom_app password=P@ssW0rd sslmode=disable"
    expectedSystemConnStr := "host=localhost port=5432 user=root dbname=postgres password=P@ssW0rd sslmode=disable"
    actualSystemConnStr, dbName := createSystemDbConnStr(Postgres, &connStr)
    assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
    assert.Equal(t, "custom_app", dbName)

    // test when db name like hostname
	connStr = "host=mysuperapp.com port=5432 user=mysuperapp dbname=mysuperapp password=123456 sslmode=disable"
	expectedSystemConnStr = "host=mysuperapp.com port=5432 user=mysuperapp dbname=postgres password=123456 sslmode=disable"
	actualSystemConnStr, dbName = createSystemDbConnStr(Postgres, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "mysuperapp", dbName)
}

func TestCreateMssqlSystemDbConnectionString(t *testing.T) {
	connStr := "sqlserver://sa:123@192.168.10.100:1433?database=custom_app"
	expectedSystemConnStr := "sqlserver://sa:123@192.168.10.100:1433?database=master"
	actualSystemConnStr, dbName := createSystemDbConnStr(Mssql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "custom_app", dbName)

	// test when db name like hostname
	connStr = "sqlserver://mysupeapp:123@mysuperapp.com:1433?database=mysuperapp"
	expectedSystemConnStr = "sqlserver://mysupeapp:123@mysuperapp.com:1433?database=master"
	actualSystemConnStr, dbName = createSystemDbConnStr(Mssql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "mysuperapp", dbName)
}

func TestCreateMysqlSystemDbConnectionString(t *testing.T) {
	connStr := "root:P@ssW0rd@tcp(127.0.0.1:3306)/custom_app?charset=utf8mb4&parseTime=True&loc=Local"
	expectedSystemConnStr := "root:P@ssW0rd@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	actualSystemConnStr, dbName := createSystemDbConnStr(Mysql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "custom_app", dbName)

	// test when db name like hostname
	connStr = "mysuperapp:P@ssW0rd@tcp(mysuperapp.com:3306)/mysuperapp?charset=utf8mb4&parseTime=True&loc=Local"
	expectedSystemConnStr = "mysuperapp:P@ssW0rd@tcp(mysuperapp.com:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	actualSystemConnStr, dbName = createSystemDbConnStr(Mysql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "mysuperapp", dbName)
}

// ####################################################################################################################

// ################################################# internal functions ###############################################
func testOpenDbWithCreateAndCheck(t *testing.T, connStr string, dialect SqlDialect) {
	db := OpenDb2(dialect, connStr, true)
	assert.NotNil(t, db)
	// Close
	CloseDb(db)
	// Drop
	DropDb(dialect, connStr)
	// Check
	checkResult := CheckDb(dialect, connStr)
	assert.Equal(t, false, checkResult)
}