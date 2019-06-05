package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func Initialize(addr, user, password, dbName string, maxOpen int) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, addr, dbName)
	gdb, err := gorm.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	gdb.DB().SetMaxOpenConns(maxOpen)
	gdb.DB().SetMaxIdleConns(0)
	gdb.DB().SetConnMaxLifetime(time.Second)

	return gdb, err
}
