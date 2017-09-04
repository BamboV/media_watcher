//controllers/gorm.go
package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	// short name for revel
	r "github.com/revel/revel"
	"github.com/bamboV/media_watcher/app/models"
	"database/sql"
)

type GormController struct {
	*r.Controller
	Txn *gorm.DB
}

// it can be used for jobs
var Gdb *gorm.DB

// init db
func InitDB() {
	var err error
	// open db
	Gdb, err = gorm.Open("postgres", "host=postgres user=media_watcher dbname=media_watcher sslmode=disable password=media_watcher")
	if err != nil {
		r.ERROR.Println("FATAL", err)
		panic( err )
	}
	Gdb.AutoMigrate(&models.Media{})
	// uniquie index if need
	//Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}


// transactions

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() r.Result {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}