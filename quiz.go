package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"goFinal/Factory"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer string
}

func parseLines(lines [][]string) []problem  {
	ret := make([]problem, len(lines))
	for i, line:=range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return  ret
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	fmt.Println("Hello! Choose hero attribute.")
	fmt.Println("1 - Agility, 2 - Strength")
	attribute := chooseAttribute()
	hero1 := Factory.GetHero(attribute)
	printDetails(hero1)

	fmt.Println("Answer all questions in Caps Lock!")
	fileName := flag.String("csv", "questions.csv", "questions")
	timeLimit := flag.Int("limit", 15, "time limit")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil{
		exit("Fail.")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


	correct := 0
	for i, p := range problems{
		fmt.Printf("Question #%d: %s = \n", i+1, p.question )
		answerChanel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChanel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nyou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerChanel:
			if answer == p.answer {
				correct++
			}
		}

	}

	fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
}
