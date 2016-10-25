package hyapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Students struct {
	Students []Student
	Count int
}

type Student struct {
	_Id string
	Name string
	Photo string
	Location string
	Cohort Cohort
	Social Social
	Job Job
}

type Cohort struct {
	Year string
	Season string
}

type Social struct {
	Website string
	Github string
	Twitter string
}

type Job struct {
	postion string
	location string
}

func GetStudents(apiKey string) (Students,error){
	
	apiUrl := "http://api.hackeryou.com/v1/"

	studentsCall, err := http.Get(apiUrl + "students/?key=" + apiKey)

	defer studentsCall.Body.Close()

	students := new(Students)

	dec := json.NewDecoder(studentsCall.Body)

	for dec.More() {
		var student Students
		err := dec.Decode(&student)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(student)
	}

	return *students, err
}