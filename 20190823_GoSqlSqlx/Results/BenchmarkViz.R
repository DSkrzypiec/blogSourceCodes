# ------------------------------------------------------------------------------- 
#
#   Author(s):  Damian Skrzypiec <damian.j.skrzypiec@gmail.com>
#
#   Date:       2019-09-03 13:58:12
#
#   Descr:      This script contains analysis and vizualization of go database/sql
#               and sqlx reading benchmark. Benchark was performed in go and results
#               was written off into .txt files (BenchResults.txt and 
#               BenchResults_TenCols.txt). Benchmark times are expressed in 
#               nanoseconds.
#
# ------------------------------------------------------------------------------- 

library(dplyr)
library(ggplot2)
library(knitr)

benchRes2cols <- read.table(file = "BenchResults.txt", header = TRUE, sep = ";")
benchRes10cols <- read.table(file = "BenchResults_TenCols.txt", header = TRUE, sep = ";")


# Function avgBenchTimeByRows prepares data.frame with average of execution time 
# of benchmark for each "NRows" seperatly based on benchmark raw data.frame. 
# Argument den transforms nanaseconds to another unit, as default to seconds.
avgBenchTimeByRows <- function(benchData, den = 10^9) {
    benchData %>%
        dplyr::group_by(NRows) %>%
        dplyr::summarise(
            AvgStdExecTime = mean(SqlStdExecTime) / den,
            AvgXExecTime = mean(SqlXExecTime) / den
            ) %>%
            dplyr::mutate(DiffPerc = (AvgXExecTime - AvgStdExecTime) / AvgXExecTime) %>%
            as.data.frame %>%
            dplyr::arrange(NRows) %>%
            as.data.frame
}

statTable2Cols <- avgBenchTimeByRows(benchRes2cols)
statTable10Cols <- avgBenchTimeByRows(benchRes10cols)

knitr::kable(statTable2Cols)
knitr::kable(statTable10Cols)

# Extend benchmark results by precentage difference: (t(sqlx) - t(std)) / t(sqlx).
benchRes2cols$DiffPerc <- (benchRes2cols$SqlXExecTime - 
                            benchRes2cols$SqlStdExecTime) / benchRes2cols$SqlXExecTime 

benchRes10cols$DiffPerc <- (benchRes10cols$SqlXExecTime - 
                            benchRes10cols$SqlStdExecTime) / benchRes10cols$SqlXExecTime 


benchRes2cols %>%
    dplyr::group_by(StdFirst, NRows) %>%
    dplyr::summarise(DiffPercAvg = mean(DiffPerc)) %>%
        as.data.frame %>%
        dplyr::arrange(StdFirst, NRows) %>%
        as.data.frame

benchRes10cols %>%
    dplyr::group_by(NRows) %>%
    dplyr::summarise(DiffPercAvg = mean(DiffPerc)) %>%
        as.data.frame %>%
        dplyr::arrange(NRows) %>%
        as.data.frame


# Function boxplotDiffPercByRows based on data.frame with benchmark results and
# chart subtitle prepares boxplot of percentage difference sqlx - sql execution 
# time.
boxplotDiffPercByRows <- function(benchData, sub) {
    x <- sort(unique(benchData$NRows))
    benchData$StdFirstName <- paste0("database/sql first in benchmark = ", 
                                     as.logical(benchData$StdFirst))

    g <- ggplot2::ggplot(data = benchData) +
            ggplot2::geom_boxplot(mapping = aes(x = factor(NRows), 
                                  y = DiffPerc), col = "black",
                                  fill ="#E0AB8A", alpha = 0.85) +
            ggplot2::ggtitle("Boxplot of precentage difference by number of rows", 
                             subtitle = sub) +
            ggplot2::xlab("Number of rows in single read") +
            ggplot2::ylab("Percentage Difference") + 
            ggplot2::theme_bw() +
            ggplot2::theme(
                plot.title = element_text(face = "bold")
            ) +
            ggplot2::scale_x_discrete(breaks = x, labels = scales::comma(x)) +
            ggplot2::facet_wrap(~StdFirstName, nrow = 2)

    return(g)
}

benchPlot2Cols <- boxplotDiffPercByRows(benchRes2cols, "Two (2) columns query")
ggplot2::ggsave(filename = "BenchmarkResults_2columns.jpg", plot = benchPlot2Cols, 
                width = 10, height = 8, units = "in")

benchPlot10Cols <- boxplotDiffPercByRows(benchRes10cols, "Ten (10) columns query")
ggplot2::ggsave(filename = "BenchmarkResults_10columns.jpg", plot = benchPlot10Cols, 
                width = 10, height = 8, units = "in")
