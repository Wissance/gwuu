package gorm_test

import (
	"github.com/stretchr/testify/assert"
	g "github.com/wissance/gwuu/gorm"
	"gorm.io/gorm"

	//"gorm.io/gorm"
	"testing"
)

const dbUser = "developer"
const dbPassword = "123"

// ############################################# test options data ####################################################
var postgresCollation = g.Collation{Encoding: "UTF8", Parameters: map[string]string{"LC_COLLATE": "C",
	"LC_CTYPE": "C", "TEMPLATE": "template0"}}

//Parameters: map[string]string{"LC_CTYPE": "en_US.utf8", "LC_COLLATE": "en_US.utf8"}

var mysqlCollation = g.Collation{Encoding: "utf8mb4", Parameters: map[string]string{"COLLATE": "utf8mb4_unicode_ci"}}
var mssqlCollation = g.Collation{Encoding: "Cyrillic_General_BIN2", Parameters: map[string]string{}}

// ####################################################################################################################

// ########################################### public functions tests #################################################

// test Build connection string

func TestBuildPostgresConnectionString(t *testing.T) {
	connStr := g.BuildConnectionString(g.Postgres, "localhost", 5432, "custom_app",
		"root", "P@ssW0rd", "disable")
	expectedConnStr := "host=localhost port=5432 user=root dbname=custom_app password=P@ssW0rd sslmode=disable"
	assert.Equal(t, expectedConnStr, connStr)
}

func TestBuildMysqlConnectionString(t *testing.T) {
	connStr := g.BuildConnectionString(g.Mysql, "127.0.0.1", 3306, "custom_app",
		"root", "P@ssW0rd", "")
	expectedConnStr := "root:P@ssW0rd@tcp(127.0.0.1:3306)/custom_app?charset=utf8mb4&parseTime=True&loc=Local"
	assert.Equal(t, expectedConnStr, connStr)
}

func TestBuildMssqlConnectionString(t *testing.T) {
	connStr := g.BuildConnectionString(g.Mssql, "192.168.10.100", 1433, "custom_app",
		"sa", "123", "")
	expectedConnStr := "sqlserver://sa:123@192.168.10.100:1433?database=custom_app"
	assert.Equal(t, expectedConnStr, connStr)
}

// test open db (system db without create)

func TestPostgresOpenSystemDb(t *testing.T) {
	cfg := gorm.Config{}
	db := g.OpenDb(g.Postgres, "127.0.0.1", 5432, "postgres", dbUser, dbPassword, "disable", false, false, &cfg,
		&postgresCollation)
	assert.NotNil(t, db)
	g.CloseDb(db)
}

func TestMysqlOpenSystemDb(t *testing.T) {
	cfg := gorm.Config{}
	db := g.OpenDb(g.Mysql, "localhost", 3306, "mysql", dbUser, dbPassword, "", false, false, &cfg,
		&mysqlCollation)
	assert.NotNil(t, db)
	g.CloseDb(db)
}

func TestMssqlOpenSystemDb(t *testing.T) {
	cfg := gorm.Config{}
	db := g.OpenDb(g.Mssql, "localhost", 1433, "master", dbUser, dbPassword, "", false, false, &cfg,
		&mssqlCollation)
	assert.NotNil(t, db)
	g.CloseDb(db)
}

// test open db with create
func TestPostgresOpenDbWithCreate(t *testing.T) {
	// Create Db when open
	cfg := gorm.Config{}
	// with collation
	connStr := g.BuildConnectionString(g.Postgres, "127.0.0.1", 5432, "pg_gwuu_examples", dbUser, dbPassword, "disable")
	testOpenDbWithCreateAndCheck(t, connStr, g.Postgres, &cfg, &postgresCollation)
	// without collation
	connStr = g.BuildConnectionString(g.Postgres, "127.0.0.1", 5432, "pg_gwuu_examples_2", dbUser, dbPassword, "disable")
	testOpenDbWithCreateAndCheck(t, connStr, g.Postgres, &cfg, nil)
}

func TestMysqlOpenDbWithCreate(t *testing.T) {
	cfg := gorm.Config{}
	// with collation
	connStr := g.BuildConnectionString(g.Mysql, "127.0.0.1", 3306, "my_gwuu_examples", dbUser, dbPassword, "")
	testOpenDbWithCreateAndCheck(t, connStr, g.Mysql, &cfg, &mysqlCollation)
	// without collation
	connStr = g.BuildConnectionString(g.Mysql, "127.0.0.1", 3306, "my_gwuu_examples_2", dbUser, dbPassword, "")
	testOpenDbWithCreateAndCheck(t, connStr, g.Mysql, &cfg, nil)
}

func TestMssqlOpenDbWithCreate(t *testing.T) {
	cfg := gorm.Config{}
	// with collation
	connStr := g.BuildConnectionString(g.Mssql, "localhost", 1433, "MsGwuuExamples", dbUser, dbPassword, "")
	testOpenDbWithCreateAndCheck(t, connStr, g.Mssql, &cfg, &mssqlCollation)
	// without collation
	connStr = g.BuildConnectionString(g.Mssql, "localhost", 1433, "MsGwuuExamples2", dbUser, dbPassword, "")
	testOpenDbWithCreateAndCheck(t, connStr, g.Mssql, &cfg, nil)
}

func TestCreateRandomDb(t *testing.T) {
	cfg := gorm.Config{}
	db, connStr := g.CreateRandomDb(g.Postgres, "127.0.0.1", 5432, dbUser, dbPassword, "disable", &cfg,
		&postgresCollation)
	assert.NotNil(t, db)
	assert.NotEmpty(t, connStr)
	check := g.CheckDb(g.Postgres, connStr, &cfg)
	assert.True(t, check)
	g.DropDb(g.Postgres, connStr, &cfg)
}

// ####################################################################################################################

// ########################################### private functions tests ################################################

func TestCreatePostgresSystemDbConnectionString(t *testing.T) {
	connStr := "host=localhost port=5432 user=root dbname=custom_app password=P@ssW0rd sslmode=disable"
	expectedSystemConnStr := "host=localhost port=5432 user=root dbname=postgres password=P@ssW0rd sslmode=disable"
	actualSystemConnStr, dbName := g.CreateSystemDbConnStr(g.Postgres, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "custom_app", dbName)

	// test when db name like hostname
	connStr = "host=mysuperapp.com port=5432 user=mysuperapp dbname=mysuperapp password=123456 sslmode=disable"
	expectedSystemConnStr = "host=mysuperapp.com port=5432 user=mysuperapp dbname=postgres password=123456 sslmode=disable"
	actualSystemConnStr, dbName = g.CreateSystemDbConnStr(g.Postgres, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "mysuperapp", dbName)
}

func TestCreateMssqlSystemDbConnectionString(t *testing.T) {
	connStr := "sqlserver://sa:123@192.168.10.100:1433?database=custom_app"
	expectedSystemConnStr := "sqlserver://sa:123@192.168.10.100:1433?database=master"
	actualSystemConnStr, dbName := g.CreateSystemDbConnStr(g.Mssql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "custom_app", dbName)

	// test when db name like hostname
	connStr = "sqlserver://mysupeapp:123@mysuperapp.com:1433?database=mysuperapp"
	expectedSystemConnStr = "sqlserver://mysupeapp:123@mysuperapp.com:1433?database=master"
	actualSystemConnStr, dbName = g.CreateSystemDbConnStr(g.Mssql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "mysuperapp", dbName)
}

func TestCreateMysqlSystemDbConnectionString(t *testing.T) {
	connStr := "root:P@ssW0rd@tcp(127.0.0.1:3306)/custom_app?charset=utf8mb4&parseTime=True&loc=Local"
	expectedSystemConnStr := "root:P@ssW0rd@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	actualSystemConnStr, dbName := g.CreateSystemDbConnStr(g.Mysql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "custom_app", dbName)

	// test when db name like hostname
	connStr = "mysuperapp:P@ssW0rd@tcp(mysuperapp.com:3306)/mysuperapp?charset=utf8mb4&parseTime=True&loc=Local"
	expectedSystemConnStr = "mysuperapp:P@ssW0rd@tcp(mysuperapp.com:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	actualSystemConnStr, dbName = g.CreateSystemDbConnStr(g.Mysql, &connStr)
	assert.Equal(t, expectedSystemConnStr, actualSystemConnStr)
	assert.Equal(t, "mysuperapp", dbName)
}

// ####################################################################################################################

// ################################################# internal functions ###############################################

func testOpenDbWithCreateAndCheck(t *testing.T, connStr string, dialect g.SqlDialect, options *gorm.Config, collation *g.Collation) {
	db := g.OpenDb2(dialect, connStr, true, true, options, collation)
	assert.NotNil(t, db)
	// Close
	g.CloseDb(db)
	// Drop
	g.DropDb(dialect, connStr, options)
	// Check
	checkResult := g.CheckDb(dialect, connStr, options)
	assert.Equal(t, false, checkResult)
}

// ####################################################################################################################
