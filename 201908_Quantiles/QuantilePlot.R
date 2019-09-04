#------------------------------------------------------------------------------- 
#
#   Author(s):  Damian Skrzypiec (damian.j.skrzypiec@gmail.com)
#
#   Date:       Tue 20 Aug 2019 09:07:13 PM CEST
#
#   Descr:      This R script produces and saves ggplot2 chart representing
#               interpolation function described in post on
#               https://dskrzypiec.github.io/blog/quantiles.
#
#------------------------------------------------------------------------------- 


library(ggplot2)


salary <- c(70, 80, 100, 140, 200)
qs <- seq(0, 1, by = 0.01)
q <- as.numeric(quantile(x = salary, probs = qs, type = 7))
qDF <- data.frame(Prob = qs, Quantile = q)

qPlot <- ggplot2::ggplot(data = qDF, aes(x = Prob)) +
            ggplot2::geom_hline(yintercept = salary, linetype = "dashed") + 
            ggplot2::geom_point(aes(y = Quantile), size = 2, col = "#DC6900") +
            ggplot2::ggtitle("Quantile plot for salary sample") +
            ggplot2::xlab("Probability - p") +
            ggplot2::ylab("Quantile") +
            ggplot2::scale_y_continuous(breaks = c(seq(70, 200, by = 25), 200)) + 
            ggplot2::theme_bw()

ggplot2::ggsave("SalaryQuantilePlot.jpg", qPlot, width = 10, height = 7, units = "in")
