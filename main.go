package main

import (
	"fmt"
	"net/http"
	"os"
	
)

// GetGreeting is exported but has no doc comment -> flagged by "revive"
func GetGreeting(name string) string {
	greeting := fmt.Sprintf("%d", name) // wrong format verb for a string -> flagged by "govet"
	fmt.Println("greeting",greeting)
	return greeting
}

func checkFlag(x bool) {
	if x == true { // should just be "if x" -> flagged by "gosimple"
		fmt.Println("flag is true")
	}
}

func fetchData(url string) {
	resp, _ := http.Get(url) // response body never closed -> flagged by "bodyclose"
	fmt.Println(resp.Status)
}

func writeFile() {
	f, _ := os.Create("output.txt")
	f.Close() // error from Close() ignored -> flagged by "errcheck"

	x := 5
	x = 10 // first assignment to x is never used -> flagged by "ineffassign"
	fmt.Println(x)
}

func buildQuery(userInput string) string {
	// naive string concatenation for a query -> flagged by "gosec" (G201/SQL injection pattern)
	query := "SELECT * FROM users WHERE name = '" + userInput + "'"
	return query
}

func main() {
	fmt.Println(GetGreeting("world"))
	checkFlag(true)
	fetchData("https://example.com")
	writeFile()
	fmt.Println(buildQuery("arun"))
}
