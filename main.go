package main 

import (
	"fmt"
	"./lib"
)

func main() {
	apiKey := "$2a$10$7/WQI00TjvOTlbyBHsRRU..W3n.gVSZ8uPOipaBNHogFv9AjRF6jy"

	fmt.Println("Hy Api")
	res, err := hyapi.GetStudents(apiKey)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}