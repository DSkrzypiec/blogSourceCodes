package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "**********"
	dbname   = "Random"
)

func ReadSQLStd(dbConn *sql.DB, nrows int) {
	rows, qErr := dbConn.Query(PointQuery(nrows))
	if qErr != nil {
		fmt.Println("Couldn't execute the query")
	}
	defer rows.Close()

	obs := make([]Point, 0, nrows)
	var (
		x int
		y float64
	)

	for rows.Next() {
		rows.Scan(&x, &y)
		obs = append(obs, Point{x, y})
	}
}

type Point struct {
	X int     `db:"x1"`
	Y float64 `db:"x2"`
}

func ReadSQLX(dbConn *sqlx.DB, nrows int) {
	obs := make([]Point, 0, nrows)
	dbConn.Select(&obs, PointQuery(nrows))
}

func PointQuery(nrows int) string {
	q := fmt.Sprintf(`
		SELECT X1, X2 FROM dbo.forbench LIMIT %d 
	`, nrows)
	return q
}

func PGConnString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}
