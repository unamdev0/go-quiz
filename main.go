package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
			return nil, fmt.Errorf("error in reading data from CSV; fileName %s\n error %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening the file;fileName %s\n error %s", fileName, err.Error())
	}

}

func parseProblem(lines [][]string) []problem {

	problems := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		answer, err := strconv.Atoi(lines[i][1])
		if err != nil {
			panic(err)
		}
		problems[i] = problem{question: lines[i][0], answer: answer}
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

	correctAnswer := 0

	tObj := time.NewTimer(time.Duration(timer) * time.Second)

	// using channel cause program can will exit after 30 seconds if no answer is given
	// if the user inputs answer we can access it using channel. On line 70 go routine is
	// called so that we don't have to wait for input from user's end
	ansC := make(chan int)

problemLoop:
	for i, p := range problems {
		var answer int
		fmt.Printf("Problem %d : %s : ", i+1, p.question)

		go func() {
			fmt.Scanf("%d", &answer)
			ansC <- answer
		}()

		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop

		case iAns := <-ansC:
			if iAns == p.answer {
				correctAnswer++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}

	fmt.Printf("Your result is %d out of %d", correctAnswer, len(problems))

}

func exit(mes string) {
	fmt.Printf("%s", mes)
	os.Exit(1)
}
