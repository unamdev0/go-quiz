package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   int
}

func problemPuller(fileName string) ([]problem, error) {

	if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)
		if lines, err := csvR.ReadAll(); err == nil {
			problems := parseProblem(lines)
			return problems, nil
		} else {
			return nil, fmt.Errorf("Error in reading data from CSV; fileName %s\n error %s\n", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("Error in opening the file;fileName %s\n error %s\n", fileName, err.Error())
	}

}

func parseProblem(lines [][]string) []problem {

	problems := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		problems[i] = problem{question: lines[i][0], answer: lines[i][1]}
	}

	return problems

}

func main() {

	fileName := "quiz.csv"
	timer := 30

	problems, err := problemPuller(fileName)

	if err != nil {
		exit(fmt.Sprintf("something went wrong: %s", err.Error()))
	}

	correctAnswer = 0

	tObj := time.NewTimer(time.Duration(timer) * time.Second)
	ansC = make(chan string)

}


func exit(mes string){
	fmt.Printf(mes)
	os.Exit(1)
}