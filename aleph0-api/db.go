package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
   "log"
)

func dbsetup(){
  db, err := gorm.Open("postgres", "user=railerde dbname=aleph0 sslmode=disable password=liberatorarchon7!4$K")
  if err != nil {
    log.Fatal(err)
  }
// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

// You can also use an existing database connection handle
// dbSql, _ := sql.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
// db := gorm.Open("postgres", dbSql)

// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
  db.DB()

  // Then you could invoke `*sql.DB`'s functions with it
  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)

// Disable table name's pluralization
//  db.SingularTable(true)
//  db.DropTable(&Post{})
//  db.DropTable(&Reply{})

//  db.DropTable(&Apikey{})
  db.CreateTable(&Apikey{})
  db.AutoMigrate(&Apikey{})
//  db.CreateTable(&Post{})
//  db.AutoMigrate(&Post{})
//  db.CreateTable(&Reply{})
//  db.AutoMigrate(&Reply{})
  db.CreateTable(&User{})
  db.AutoMigrate(&User{})
  db.CreateTable(&Item{})
  db.AutoMigrate(&Item{})
//  db.DropTable(&Contact{})
  db.CreateTable(&Contact{})
  db.AutoMigrate(&Contact{})
  db.CreateTable(&Signup{})
  db.AutoMigrate(&Signup{})
}
