package hyapi

import ( 
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

type Operations struct {
	Operations []Operation
	Count int `json:"count"`
}

type Operation struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Social Social
}

type Instructors struct {
	Instructors []Instructor
	Count int `json:"count"`
}

type Instructor struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Social Social
}


const apiUrl string = "http://api.hackeryou.com/v1/"

func GetStudents(apiKey string) (Students,error){
	

	studentsCall, err := http.Get(apiUrl + "students/?key=" + apiKey)

	defer studentsCall.Body.Close()

	dec := json.NewDecoder(studentsCall.Body)

	//create var for items to be in 
	var student Students
	//Loop through json stream
	for dec.More() {
		//Decode each
		err := dec.Decode(&student)
		if err != nil {
			return student, err
		}
	}

	return student, err
}

func GetOperations(apiKey string) (Operations, error) {
	opsCall, err := http.Get(apiUrl + "operations/?key=" + apiKey)

	defer opsCall.Body.Close()

	dec := json.NewDecoder(opsCall.Body)

	var operations Operations

	for dec.More() {
		err := dec.Decode(&operations)
		if err != nil {
			return operations, err
		}
	}

	return operations, err
}

func GetInstructors(apiKey string) (Instructors, error) {
	instructorsCall, err := http.Get(apiUrl + "instructors/?key=" + apiKey) 

	defer instructorsCall.Body.Close()

	dec := json.NewDecoder(instructorsCall.Body)

	var instructors Instructors

	for dec.More() {
		err := dec.Decode(&instructors)
		if err != nil {
			return instructors, err
		}
	}

	return instructors, err
}