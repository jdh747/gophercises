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
)

func main() {
	// Initialise flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	// limit := flag.Int("limit", 30, "the limit for the quiz in seconds")
	flag.Parse()

	problems := parseProblems(csvFilename)
	runQuiz(problems)
}

func parseProblems(csvFilename *string) []problem {
	file, _ := os.Open(*csvFilename)
	csvReader := csv.NewReader(file)

	var problems []problem

	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, problem{
			strings.TrimSpace(line[0]),
			strings.TrimSpace(line[1]),
		})
	}
	return problems
}

func runQuiz(problems []problem) {
	correct := 0
	inputReader := bufio.NewReader(os.Stdin)
	for i, prob := range problems {
		fmt.Printf("Question %d) %s: ", i+1, prob.question)
		result, _ := inputReader.ReadString('\n')
		if strings.TrimSpace(result) == prob.answer {
			correct++
		}
	}
	fmt.Printf("Complete. %d/%d correct\n", correct, len(problems))
}

type problem struct {
	question string
	answer   string
}
