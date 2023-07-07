package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

//Questions and answers
type QandA struct {
	Question string
	Anwser   string
}

func main() {

	filename := flag.String("file", "problems.csv", "CSV filename")
	duration := flag.Duration("timer", 30*time.Second, "time limit")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	questions := make([]QandA, 0)

	for _, record := range data {
		if len(record) >= 2 {
			q := QandA{
				Question: record[0],
				Anwser:   record[1],
			}
			questions = append(questions, q)
		}
	}

	var (
		correct int
		input   string
	)

	fmt.Println("Hi, this is quiz game. Anwser the questions pls\nPress enter when You are ready")
	fmt.Scanln()
	timer := time.NewTimer(*duration)

	for i, record := range questions {
		fmt.Printf("#%d  %s is equal: \n", i+1, record.Question)
		answerCh := make(chan string)
		go func() {
			fmt.Scanln(&input)
			answerCh <- input
		}()
		select {
		case <-timer.C:
			fmt.Printf("time's up\n_______\nYour score: %v/%v", correct, len(questions))
			return
		case input := <-answerCh:
			if input == record.Anwser {
				fmt.Println("correct, noice")
				correct++
			} else {
				fmt.Println("sorry, wrong anwser")
			}
		}
	}

	fmt.Printf("_______\nYour score: %v/%v", correct, len(questions))
}
