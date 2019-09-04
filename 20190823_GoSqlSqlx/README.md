# Go database/sql vs jmoiron/sqlx reading benchmark source codes

This repo contains source codes of performing benchmark between go
`database/sql` package and `jmoiron/sqlx` package. This benchmark was described
in [https://dskrzypiec.dev/gosqlstdsqlx/](https://dskrzypiec.dev/gosqlstdsqlx/)
post. Structure of codes is as follows:

* `DataInsert` contains single python script for inserting 10mil records into
    PostgreSQL database.
* `GoSqlSqlxBench` contains go implementation of reading from the database using
  `database/sql` and `jmoiron/sqlx` package. Main program performs benchmark and
    writes results into `.txt` files.
* `Results` contains files with benchmark results and `BenchmarkViz.R` R script
    for producing statistics and charts.

Raw results of benchmarks are stored in files `Results/BenchResults.txt` for two
columns case and in `Results/BenchResults_TenCols.txt` for ten columns case.
