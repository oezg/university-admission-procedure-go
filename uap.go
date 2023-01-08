package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	numberExams = 3
	threshold   = 60
)

var (
	numberApplicants int
	numberAdmission  int
	meanScore        float64
	decision         bool
	applicants       []Applicant
)

type Applicant struct {
	firstName string
	lastName  string
	gpa       float64
}

func main() {
	getApplicants()
	showAdmitted()
}

func showAdmitted() {
	fmt.Println("Successful applicants:")
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}
		return applicants[i].lastName < applicants[j].lastName
	})
	for i := 0; i < numberAdmission; i++ {
		fmt.Printf("%s %s\n", applicants[i].firstName, applicants[i].lastName)
	}
}

func getApplicants() {
	fmt.Scanln(&numberApplicants)
	fmt.Scanln(&numberAdmission)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < numberApplicants; i++ {
		scanner.Scan()
		fields := strings.SplitN(scanner.Text(), " ", 3)
		point, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		newApplicant := Applicant{
			firstName: fields[0],
			lastName:  fields[1],
			gpa:       point,
		}
		applicants = append(applicants, newApplicant)
	}
}

func showResult() {
	if decision {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}

func getMeanScore() {
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
}

func getDecision() {
	fmt.Println(meanScore)
	decision = meanScore >= threshold
}
