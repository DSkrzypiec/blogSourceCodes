package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func BenchmarkSQLStd(b *testing.B) {
	db, err := sql.Open("postgres", PGConnString())
	if err != nil {
		fmt.Println("Cannot open database :(")
	}

	for i := 0; i < b.N; i++ {
		ReadSQLStd(db, 10000)
	}
}

func BenchmarkSQLX(b *testing.B) {
	db, err := sqlx.Connect("postgres", PGConnString())
	if err != nil {
		fmt.Println("Cannot open database :(")
	}

	for i := 0; i < b.N; i++ {
		ReadSQLX(db, 10000)
	}
}
