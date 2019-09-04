package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
)

type StdReader func(*sql.DB, int)
type SqlxReader func(*sqlx.DB, int)

// QueryTimeStats keeps time results of executing SQL standard and sqlx
// readings. Execution time is stored in nanoseconds.
type QueryTimeStats struct {
	Id             int
	StdFirst       bool
	SqlStdExecTime int64
	SqlXExecTime   int64
	NRows          int
}

// String method for print.
func (q QueryTimeStats) String() string {
	qStr := fmt.Sprintf("Id = %d | StdFirst = %t | StdTime = %d | XTime = %d | N = %d \n",
		q.Id, q.StdFirst, q.SqlStdExecTime, q.SqlXExecTime, q.NRows)
	return qStr
}

// TxtString method returns string which is used while reading results into file
// with ";" separator.
func (q QueryTimeStats) TxtString() string {
	var sf int = 0
	if q.StdFirst {
		sf = 1
	}

	qs := fmt.Sprintf("%d;%d;%d;%d", sf, q.SqlStdExecTime, q.SqlXExecTime,
		q.NRows)
	return qs
}

// Function MeasureQueries measures single execution of standard and sqlx
// reading.
func MeasureQueries(fStd StdReader, fx SqlxReader, db *sql.DB,
	dbx *sqlx.DB, nrows, id int) QueryTimeStats {
	stdFirst := rand.Float64() > 0.50

	if stdFirst {
		startStd := time.Now()
		fStd(db, nrows)
		endStd := time.Now()
		timeStdNS := endStd.Sub(startStd).Nanoseconds()

		startX := time.Now()
		fx(dbx, nrows)
		endX := time.Now()
		timeXNS := endX.Sub(startX).Nanoseconds()

		return QueryTimeStats{id, true, timeStdNS, timeXNS, nrows}
	}

	startX := time.Now()
	fx(dbx, nrows)
	endX := time.Now()
	timeXNS := endX.Sub(startX).Nanoseconds()

	startStd := time.Now()
	fStd(db, nrows)
	endStd := time.Now()
	timeStdNS := endStd.Sub(startStd).Nanoseconds()

	return QueryTimeStats{id, false, timeStdNS, timeXNS, nrows}
}

func BenchmarkReadings(fStd StdReader, fx SqlxReader, db *sql.DB,
	dbx *sqlx.DB, n int, w *bufio.Writer, nrows []int, sleepTime int) {

	for _, nrow := range nrows {
		for i := 0; i < n; i++ {
			res := MeasureQueries(fStd, fx, db, dbx, nrow, i)
			if i < 5 {
				continue
			}
			fmt.Print(res)
			tmp := fmt.Sprintf("%s;%d\n", res.TxtString(), sleepTime)
			w.WriteString(tmp)
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}
}
