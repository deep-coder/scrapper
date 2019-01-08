# scrapper

default Args has been set.
If we want to pass some other args we need to run

go run main.go hiverhq.com https://hiverhq.com/ 5

else just

go run main.go

Logic:

scrapping html text nodes and putting in worker thread and putting all words as keys in map and word count as values

Then building Max Heap and picking the top 5 most used words.