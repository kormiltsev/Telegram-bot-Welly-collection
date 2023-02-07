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

// save database url
func SetURL(url string) {
	PGurl.URL = url
}

// UploadRowPostgres push new row to table
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
	return nil
}
