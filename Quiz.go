package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type questions struct {
	q string
	a string
}

func main() {

	csvfile := flag.String("csv", "temp.cs", "csv file has to be loaded")
	limit := flag.Int("limit", 10, "set the time limit for the quiz in seconds")
	flag.Parse()
	file, _ := os.Open(*csvfile)
	reader := csv.NewReader(file)
	// reader.Comma = '.'
	data, _ := reader.ReadAll()
	problems := make([]questions, len(data))
	correct := 0
	timeLimit := time.NewTimer(time.Duration(*limit) * time.Second)
	ansChan := make(chan string)
quizLoop:
	for i, row := range data {
		fmt.Printf("Question # %d %s\n", i+1, row[0])
		problems[i] = questions{
			q: row[0],
			a: strings.TrimSpace(row[1]),
		}
		go func() {
			var ans string
			fmt.Scan(&ans)
			ansChan <- ans
		}()
		select {
		case <-timeLimit.C:
			fmt.Println("Times Up!!")
			break quizLoop
		case ans := <-ansChan:

			if ans == problems[i].a {
				correct++
				fmt.Println("You are Correct!")
			} else {
				fmt.Println("Nope!!")
			}
		}

	}

	
	fmt.Println("your total marks :", correct, "out of", len(problems))

}
