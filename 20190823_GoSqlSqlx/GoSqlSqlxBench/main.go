package main

import (
	"bufio"
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const headers = "StdFirst;SqlStdExecTime;SqlXExecTime;NRows;SleepTime\n"

func main() {
	fileName := flag.String("fileName", "BenchResults2Cols.txt", "Output file with 2 cols results")
	fileName2 := flag.String("fileName2", "BenchResults10Cols.txt", "Output file with 10 cols results")
	sleepTime := flag.Int("sleepTime", 2, "Sleep time in seconds")
	flag.Parse()

	db, dbx := OpenDBConns()
	w, wTen := CreateBufWriters(*fileName, *fileName2)
	rows := []int{1000, 10000, 100000, 1000000, 10000000}

	BenchmarkReadings(ReadSQLStd, ReadSQLX, db, dbx, 100, w, rows, *sleepTime)
	w.Flush()

	BenchmarkReadings(ReadSQLStdTenCols, ReadSQLXTenCols, db, dbx, 100, wTen, rows, *sleepTime)
	wTen.Flush()
}

// This function returns active connection to the same PostrgeSQL database. One
// connection is in form standrad sql.DB and other one is for sqlx package
// purposes.
func OpenDBConns() (*sql.DB, *sqlx.DB) {
	db, err := sql.Open("postgres", PGConnString())
	if err != nil {
		log.Panic("Cannot open database :(")
	}

	dbx, err := sqlx.Connect("postgres", PGConnString())
	if err != nil {
		log.Panic("Cannot open database :(")
	}
	return db, dbx
}

// Helper function which for given names creates new files and returns
// bufio.Writer based on those files. It also writes headers on the table with
// benchmark results.
func CreateBufWriters(fName1, fName2 string) (*bufio.Writer, *bufio.Writer) {
	f, err := os.Create(fName1)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(headers)

	f2, err := os.Create(fName2)
	if err != nil {
		log.Fatal(err)
	}
	w2 := bufio.NewWriter(f2)
	w2.WriteString(headers)

	return w, w2
}
