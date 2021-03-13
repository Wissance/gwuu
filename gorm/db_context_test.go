package gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// public functions tests

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

}

func TestMysqlOpenSystemDb(t *testing.T) {

}

func TestMssqlOpenSystemDb(t *testing.T) {

}

// test open db with create

func TestPostgresOpenDbWithCreate(t *testing.T) {

}

func TestMysqlOpenDbWithCreate(t *testing.T) {

}

func TestMssqlOpenDbWithCreate(t *testing.T) {

}