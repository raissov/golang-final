package main

import (
	"fmt"
	_ "goFinal/Builder"
	"goFinal/Factory"
)

func main() {
	fmt.Println("Hello! Choose hero attribute.")
	fmt.Println("1 - Agility, 2 - Strength")
	attribute := chooseAttribute()
	hero1 := Factory.GetHero(attribute)
	printDetails(hero1)


	//fmt.Println("Answer all questions in Caps Lock!")
	//fileName := flag.String("csv", "questions.csv", "questions")
	//timeLimit := flag.Int("limit", 15, "time limit")
	//flag.Parse()
	//
	//file, err := os.Open(*fileName)
	//if err != nil{
	//	exit("Fail.")
	//}
	//r := csv.NewReader(file)
	//lines, err := r.ReadAll()
	//problems := parseLines(lines)
	//
	//timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//
	//
	//correct := 0
	//for i, p := range problems{
	//	fmt.Printf("Question #%d: %s = \n", i+1, p.question )
	//	answerChanel := make(chan string)
	//	go func() {
	//		var answer string
	//		fmt.Scanf("%s\n", &answer)
	//		answerChanel <- answer
	//	}()
	//
	//	select {
	//	case <-timer.C:
	//		fmt.Printf("\nyou scored %d out of %d.\n", correct, len(problems))
	//		return
	//	case answer := <-answerChanel:
	//		if answer == p.answer {
	//			correct++
	//		}
	//	}
	//
	//}
	//
	//fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
}

func chooseAttribute() string  {
	var attCode int
	var attribute string
	fmt.Scan(&attCode)
	if attCode < 4 && attCode > 0 {
		switch attCode {
		case 1:
			attribute = "agility"
		case 2:
			attribute = "strength"
		}
	} else {
		fmt.Println("Enter values from 1 to 3")
		chooseAttribute()
	}
	return attribute
}

func printDetails(g Factory.IHero) {
	fmt.Printf("Name: %s \n", g.GetName())
	fmt.Printf("Attribute: %s \n", g.GetAttribute())
}

