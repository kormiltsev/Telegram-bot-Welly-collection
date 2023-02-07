package product

import (
	"context"
	"log"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

type PGcon struct {
	URL string
}

var PGurl = PGcon{}

func SetURL(url string) {
	PGurl.URL = url
}

// CREATE TABLE IF NOT EXISTS urls(
//
//		userid integer not null,
//		alias text not null unique,
//		original text not null,
//		deleted bool
//	  )
//
// CREATE TABLE $1
// var sqlQtyColumns =
// SELECT COUNT(*)
// FROM INFORMATION_SCHEMA.COLUMNS
// WHERE table_catalog = $1
// AND table_name = $2
// ;
// ALTER TABLE $1 ADD COLUMN $2 text;
// var sqlAddColumn = `
// CREATE TABLE $1(

func UploadRowPostgres(w *Welly) error {
	connString := PGurl.URL
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(context.Background(), ";")
	if err != nil {
		log.Println(err)
		return err
	}

	// create table if not exists:
	var sqlCreateTable = `
	CREATE TABLE IF NOT EXISTS welly_catalog(
	id serial primary key,
	uniqid text not null,
	userid text not null,
	itemid text not null,
	manufacture text not null,
	model text not null,
	color text,
	coments text,
	title_photo text not null,
	catalog_photos text[],
	created_at TIMESTAMPTZ DEFAULT Now(),
	);
	`
	_, err = conn.Exec(context.Background(), sqlCreateTable) //, income.TableName)
	if err != nil {
		log.Println("connecting to welly table")
	}

	// insert
	var sqlInsertSlice = `
INSERT INTO welly_catalog(uniqid, userid, itemid, manufacture, model, color, coments, title_photo, catalog_photos) 
values($1, $2, $3, $4, $5, $6, $7, $8, $9)
;`
	_, err = conn.Exec(ctx, sqlInsertSlice, w.UniqID, w.UserID, w.ItemID, w.Manufacture, w.Model, w.Color, w.Comments, w.TitleFoto, w.AllFoto) //income.TableName,
	if err != nil {
		log.Printf("Unable to insert: %s\n", err)
	}

	// read
	// var id int32
	// var selectSlice []string
	// err = conn.QueryRow("select id, strings from slice_test order by id desc limit 1").Scan(&id, &selectSlice)
	// if err != nil {
	// 	log.Fatalf("Unable to select: %v", err)
	// }
	// fmt.Println("id:", id, "selectSlice", selectSlice)
	// conn.Close()
	return nil
}
