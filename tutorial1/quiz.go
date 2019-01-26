package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// Initialise flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the limit for the quiz in seconds")
	flag.Parse()

	problems := parseProblems(csvFilename)
	runQuiz(problems, *timeLimit)
}

func parseProblems(csvFilename *string) []problemT {
	file, error := os.Open(*csvFilename)
	if error != nil {
		log.Fatal(error)
	}
	csvReader := csv.NewReader(file)

	var problems []problemT

	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, problemT{
			strings.TrimSpace(line[0]),
			strings.TrimSpace(line[1]),
		})
	}
	return problems
}

func runQuiz(problems []problemT, timeLimit int) {
	correct := 0
	inputReader := bufio.NewReader(os.Stdin)
	resultChan := make(chan string, 1)
	for i, problem := range problems {
		fmt.Printf("Question %d) %s: ", i+1, problem.question)

		go func() {
			result, _ := inputReader.ReadString('\n')
			resultChan <- result
		}()

		select {
		case result := <-resultChan:
			if strings.TrimSpace(result) == problem.answer {
				correct++
			}
		case <-time.After(time.Duration(timeLimit) * time.Second):
			println("*Timeout*")
		}

	}
	fmt.Printf("Complete. %d/%d correct\n", correct, len(problems))
}

type problemT struct {
	question string
	answer   string
}
