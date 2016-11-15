package main 

import (
	"fmt"
	"./lib"
)

func main() {
	apiKey := "$2a$10$7/WQI00TjvOTlbyBHsRRU..W3n.gVSZ8uPOipaBNHogFv9AjRF6jy"

	fmt.Println("Hy Api")
	students, err := hyapi.GetStudents(apiKey)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(students)

	ops, opsErr := hyapi.GetOperations(apiKey)

	if opsErr != nil {
		fmt.Println(opsErr)
	}

	fmt.Println(ops)

	instructors, instErr := hyapi.GetInstructors(apiKey)

	if instErr != nil {
		fmt.Println(instErr)
	}

	fmt.Println(instructors)
}