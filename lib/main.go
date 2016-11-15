package hyapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Students struct {
	Students []Student
	Count int `json:"count"`
}

type Student struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	Photo string `json:"photo"`
	Location string `json:"location"`
	Cohort Cohort
	Social Social
	Job Job
}

type Cohort struct {
	Year int `json:"year"`
	Season string `json:"season"`
}

type Social struct {
	Website string `json:"website"`
	Github string `json:"github"`
	Twitter string `json:"twitter"`
}

type Job struct {
	Postion string `json:"position"`
	Location string	`json:"location"`
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