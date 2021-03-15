This is a sel of ultimate utils when using GO language to develop web applications.

Contains following tools:

* gorm - set of functions to extend gorm features:
    - build connection string (Postgres, Mssql, Mysql)
    - create database (Postgres, Mssql, Mysql)
    - drop database (Postgres, Mssql, Mysql)
    - get next identifier (sometimes GORM is unable to create entities with auto-generated identifiers therefore i have to gen it manually)
    - get portion of data

To see how to work with Create and Drop DB please see following test:

https://github.com/Wissance/gwuu/blob/master/gorm/db_context_test.go

But not all features were covered with tests, Pagaination and Get next id were experimentally tested in projects (but in near future these function will be also covered with tests). See additional functions here:

https://github.com/Wissance/gwuu/blob/master/gorm/db_utils.go
