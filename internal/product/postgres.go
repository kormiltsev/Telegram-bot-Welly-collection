package product

import (
	"github.com/go-pg/pg"
)

var db *pg.DB

func StartConnection(conf Confpg) {
	db = pg.Connect(&pg.Options{
		Addr:     conf.Adress,
		User:     conf.User,
		Password: conf.Password,
		Database: conf.DB,
	})
	CreateTable()
}

// func DbClose() {
// 	db.Close()
// }

func CreateTable() {
	// var urler Baserow
	// err := db.CreateTable(&urler, &orm.CreateTableOptions{
	// 	Temp:          false,
	// 	IfNotExists:   true,
	// 	FKConstraints: true,
	// })
	// panicIf(err)
	// //return err
}
