package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	filename         = "applicant_list.txt"
	numberPriorities = 3
	numberExams      = 5
)

var (
	maxAdmitted     int
	applicants      []Applicant
	departmentNames = []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}
	departments     = make(map[string][]Applicant)
)

type Applicant struct {
	firstName, lastName string
	scores              map[string]float64
	priorities          [numberPriorities]string
}

func main() {
	getMaxAdmitted()
	getApplicants()
	fillDepartments()
	saveDepartments()
}

func getMaxAdmitted() {
	fmt.Println("Enter the maximum number of students that will be admitted to each department:")
	_, err := fmt.Scanln(&maxAdmitted)
	if err != nil {
		log.Fatal(err)
	}
}

func getApplicants() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Could not close the file", file)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)
		examScores := make([]float64, numberExams)
		for i := 0; i < numberExams; i++ {
			examScores[i], err = strconv.ParseFloat(data[i+2], 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		applicant := Applicant{
			firstName:  data[0],
			lastName:   data[1],
			scores:     make(map[string]float64, numberExams),
			priorities: [numberPriorities]string{data[7], data[8], data[9]},
		}
		applicant.scores["Physics"] = examScores[0]
		applicant.scores["Chemistry"] = examScores[1]
		applicant.scores["Mathematics"] = examScores[2]
		applicant.scores["Engineering"] = examScores[3]
		applicant.scores["Admission"] = examScores[4]
		applicants = append(applicants, applicant)
	}
}

func fillDepartments() {
	for _, departmentName := range departmentNames {
		departments[departmentName] = []Applicant{}
	}
	for priority := 0; priority < numberPriorities; priority++ {
		for department := range departments {
			sortApplicants(applicants, department)
			admitDepartmentPriorityN(department, priority)
		}
	}
}

func sortApplicants(applicants []Applicant, department string) {
	sort.Slice(applicants, func(i, j int) bool {
		if getScore(applicants[i], department) != getScore(applicants[j], department) {
			return getScore(applicants[i], department) > getScore(applicants[j], department)
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}
		return applicants[i].lastName < applicants[j].lastName
	})
}

func getScore(applicant Applicant, department string) float64 {
	var meanScore float64
	switch department {
	case "Biotech":
		meanScore = (applicant.scores["Physics"] + applicant.scores["Chemistry"]) / 2.0
	case "Chemistry", "Mathematics":
		meanScore = applicant.scores[department]
	case "Engineering", "Physics":
		meanScore = (applicant.scores[department] + applicant.scores["Mathematics"]) / 2.0
	default:
		meanScore = 0
	}
	return math.Max(meanScore, applicant.scores["Admission"])
}

func admitDepartmentPriorityN(department string, priority int) {
	k := 0
	for _, applicant := range applicants {
		if department == applicant.priorities[priority] && len(departments[department]) < maxAdmitted {
			departments[department] = append(departments[department], applicant)
		} else {
			applicants[k] = applicant
			k++
		}
	}
	applicants = applicants[:k]
}

func saveDepartments() {
	for department, enrollment := range departments {
		filename := fmt.Sprintf("%s.txt", strings.ToLower(department))
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal("Could not close the file", file)
			}
		}(file)

		sortApplicants(enrollment, department)

		for _, candidate := range enrollment {
			score := getScore(candidate, department)
			_, err := fmt.Fprintf(file, "%s %s %.1f\n", candidate.firstName, candidate.lastName, score)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
