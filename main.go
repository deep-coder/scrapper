package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var wordCount = make(map[string]int, 1000)

func main() {
		domain := "hiverhq.com"
		website := "https://hiverhq.com/"
		count := "5"
		if len(os.Args) > 3{
			domain = os.Args[1]
			website = os.Args[2]
			count = os.Args[3]
		}

	// some validation

	contentDelivery := make(chan string, 1000)
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.MaxDepth(0),
	)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})

	c.OnHTML("h4", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})

	c.OnHTML("h5", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})

	c.OnHTML("h6", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		contentDelivery <- e.Text
	})


	c.Visit(website)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	workers(wg, contentDelivery)

	c.Wait()
	close(contentDelivery)
	wg.Wait()
	fmt.Println(wordCount)
	wordSlice := mapToSlice(wordCount)
	buildHeap(wordSlice, len(wordSlice))

	for i := range wordSlice {
		tempCount, _ := strconv.Atoi(count)
		if i >= tempCount {
			break
		}
		fmt.Println("Word:", wordSlice[0].word, " count:", wordSlice[0].count)
		// Move current root to end
		wordSlice[0], wordSlice[len(wordSlice)-1] = wordSlice[len(wordSlice)-1], wordSlice[0]
		wordSlice = wordSlice[:len(wordSlice)-1]
		// call max heapify on the reduced heap
		maxHeapify(wordSlice, len(wordSlice), 0)
	}
}

func workers(wg *sync.WaitGroup, contentDelivery chan string) {
	go func() {
		for str := range contentDelivery {
			temp := strings.Split(str, " ")
			for i := 0; i < len(temp); i++ {
				tempWord := strings.ToLower(strings.TrimSpace(temp[i]))
				wordCount[tempWord]++
			}
		}
		delete(wordCount, "")
		wg.Done()
	}()
}

type word struct {
	word  string
	count int
}

func mapToSlice(wordMap map[string]int) []word {
	wordSlice := make([]word, 0, len(wordMap))
	for k, v := range wordMap {
		wordSlice = append(wordSlice, word{word: k, count: v})
	}

	return wordSlice
}

func maxHeapify(input []word, size, rootIndex int) {
	largest := rootIndex
	left := 2*rootIndex + 1
	right := 2*rootIndex + 2

	if left < size && input[left].count > input[largest].count {
		largest = left
	}

	if right < size && input[right].count > input[largest].count {
		largest = right
	}

	if largest != rootIndex {
		input[largest], input[rootIndex] = input[rootIndex], input[largest]

		maxHeapify(input, size, largest)
	}
}

func buildHeap(input []word, size int) {
	for i := (size / 2) - 1; i >= 0; i-- {
		maxHeapify(input, size, i)
	}
}
