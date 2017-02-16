package main

import (
	"fmt"
	"net/http"
)
// Bu kodu çalıştırmak için command line'dan `go run task.go` çalıştırman yeterli.

//readBook reads the book at filePath. Keep the at a glabal variable at access it at 'count' and 'query' functions 
func readBook(filePath string) {
	//	YOUR CODE HERE. Read the book and save it to a global variable, something like `var Book [][]string`
}
func query(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("c") //Get Chapter from url as string. you have to convert it to int by using strconv.Atoi
	p := q.Get("p") //Get Paragraph from url as string, you have to convert it to int
	result := ""
	/*
		YOUR CODE HERE
	*/
	fmt.Fprint(w, result)
}
func count(w http.ResponseWriter, r *http.Request) {
	chapCount := 0
	paraCount := 0
	/*
		YOUR CODE HERE
	*/
	fmt.Fprintf(w, "chapter: %d\nparagraph: %d\n", chapCount, paraCount)
}
func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func main() {
	readBook("book.txt")
	http.HandleFunc("/count", count)
	http.HandleFunc("/query", query)
	http.HandleFunc("/", otherwise)
	http.ListenAndServe(":8080", nil)
}
