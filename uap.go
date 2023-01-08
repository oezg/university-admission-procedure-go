package main

import (
	"fmt"
	"log"
)

const (
	numberExams = 3
	threshold   = 60
)

func main() {
	meanScore := getMeanScore()
	decision := getDecision(meanScore)
	showResult(decision)
}

func showResult(decision bool) {
	if decision {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}

func getMeanScore() (meanScore float64) {
	var score float64
	var total float64
	for i := 0; i < numberExams; i++ {
		//fmt.Printf("Please enter the score of exam no. %d:\n", i+1)
		_, err := fmt.Scanln(&score)
		if err != nil {
			log.Fatal("Enter only a number")
		}
		total += score
	}
	meanScore = total / numberExams
	return
}

func getDecision(meanScore float64) (decision bool) {
	fmt.Println(meanScore)
	decision = meanScore >= threshold
	return
}
