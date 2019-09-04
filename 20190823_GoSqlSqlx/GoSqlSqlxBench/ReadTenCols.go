package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PointTen struct {
	X1  int     `db:"x1"`
	X2  float64 `db:"x2"`
	X3  string  `db:"x3"`
	X4  int     `db:"x4"`
	X5  float64 `db:"x5"`
	X6  string  `db:"x6"`
	X7  int     `db:"x7"`
	X8  float64 `db:"x8"`
	X9  string  `db:"x9"`
	X10 int     `db:"x10"`
}

func ReadSQLStdTenCols(dbConn *sql.DB, nrows int) {
	rows, qErr := dbConn.Query(TenColsQuery(nrows))
	if qErr != nil {
		fmt.Println("Couldn't execute the query")
	}
	defer rows.Close()

	obs := make([]PointTen, 0, nrows)
	var (
		x1, x4, x7, x10 int
		x2, x5, x8      float64
		x3, x6, x9      string
	)

	for rows.Next() {
		rows.Scan(&x1, &x2, &x3, &x4, &x5, &x6, &x7, &x8, &x9, &x10)
		obs = append(obs, PointTen{x1, x2, x3, x4, x5, x6, x7, x8, x9, x10})
	}
}

func ReadSQLXTenCols(dbConn *sqlx.DB, nrows int) {
	obs := make([]PointTen, 0, nrows)
	dbConn.Select(&obs, TenColsQuery(nrows))
}

func TenColsQuery(nrows int) string {
	q := fmt.Sprintf(`
		SELECT 
			X1, X2, X3, X4, X5, X6, X7, X8, X9, X10 
		FROM 
			dbo.forbench 
		LIMIT %d 
	`, nrows)
	return q
}
