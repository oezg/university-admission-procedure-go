package main

import "fmt"

const numberExams = 3

func main() {
	meanScore := getMeanScore()
	showResult(meanScore)
}

func showResult(score float64) {
	fmt.Println(score)
	fmt.Println("Congratulations, you are accepted!")
}

func getMeanScore() (meanScore float64) {
	var score float64
	var total float64
	for i := 0; i < numberExams; i++ {
		fmt.Scanln(&score)
		total += score
	}
	meanScore = total / numberExams
	return
}
