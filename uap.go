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
	filename = "applicants.txt"
)

var (
	maxAdmitted     int
	applicants      []Applicant
	departmentNames = []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}
	departments     = make(map[string][]Applicant)
)

type Applicant struct {
	firstName  string
	lastName   string
	gpa        float64
	priorities [3]string
}

func main() {
	getMaxAdmitted()
	getApplicants()
	sortDepartment()
	sortApplicants(applicants)
	fillDepartments()
	showDepartments()
}

func showDepartments() {
	for _, departmentName := range departmentNames {
		fmt.Println(departmentName)
		enrollment := departments[departmentName]
		sortApplicants(enrollment)
		for _, student := range enrollment {
			fmt.Printf("%v %v %.2f\n", student.firstName, student.lastName, student.gpa)
		}
		fmt.Println()
	}
}

func fillDepartments() {
	for i := 0; i < 3; i++ {
		admitPriorityN(i)
	}
}

func admitPriorityN(priorityRank int) {
	k := 0
	for _, applicant := range applicants {
		department := applicant.priorities[priorityRank]
		if len(departments[department]) < maxAdmitted {
			departments[department] = append(departments[department], applicant)
		} else {
			applicants[k] = applicant
			k++
		}
	}
	applicants = applicants[:k]
}

func sortApplicants(applicants []Applicant) {
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}
		if applicants[i].firstName != applicants[j].firstName {
			return applicants[i].firstName < applicants[j].firstName
		}
		return applicants[i].lastName < applicants[j].lastName
	})
}

func sortDepartment() {
	sort.Strings(departmentNames)
	for _, departmentName := range departmentNames {
		departments[departmentName] = []Applicant{}
	}
}

func getMaxAdmitted() {
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
		data := strings.SplitN(line, " ", 6)
		pointAverage, err := strconv.ParseFloat(data[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		newApplicant := Applicant{
			firstName:  data[0],
			lastName:   data[1],
			gpa:        pointAverage,
			priorities: [3]string{data[3], data[4], data[5]},
		}
		applicants = append(applicants, newApplicant)
	}
}
