library(ggplot2)

csv = read.csv("Desktop/benchmark.csv")
csv$Library = paste(csv$Language,"-",csv$Library,"-",csv$Operation,sep=" ")
csv$Records <- NULL
csv$index =  seq.int(nrow(csv))

qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(1,2, 10, 11)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nDefault JSON libraries\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))


qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(6,7, 18, 19)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nDefault Avro libraries\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))


qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(1,2, 10, 11, 6,7, 18, 19)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nAvro vs JSONSchema\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))


qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(5, 12, 13, 7, 19)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nAvro vs JSONSchema\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))



qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(1, 2, 3, 4, 10, 11, 14, 15)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nThird party JSON libs\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))



qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(6, 7, 8, 9, 18, 19)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nThird party AVRO libs\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))


qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(8, 9, 14, 15)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nFinal contestants\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))


qplot(Library, Seconds, 
      data=filter(csv, csv$index %in% c(8, 9, 14, 17)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nFinal contestants (schema validated)\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14))


qplot(Library, Seconds, 
      data=filter(csv, !csv$index %in% c(16, 12)), geom=c("line"), ylab="Seconds to process 1m records (lower=better)",
      main="\nThe Big Picture\n",
      colour=Operation, 
      label=Seconds) + 
  facet_grid(facets= Language ~ ., scales = "free_y", space = "free_y") + geom_bar(stat='identity', fill=c("white")) + 
  coord_flip() + 
  #theme_bw() + 
  geom_text(size = 3, position = position_stack(vjust = 0.5)) +
  theme(plot.title = element_text(size = 30)) +
  theme(axis.text.y=element_text(size=14)) 

?qplot
csv$is_spark = factor(csv$is_spark)
csv$tag <- factor(csv$tag, levels = csv[order(csv$questions), "tag"])

qplot(questions, tag, colour=is_spark, data=csv, c('scatter'), main='Stackoverflow questions') + theme_bw()
