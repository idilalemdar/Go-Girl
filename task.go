package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

//readBook reads the book at filePath. Keep the at a glabal variable at access it at 'count' and 'query' functions 

var Book[][]string 

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readBook(filePath string) {
	Book,err := ioutil.ReadFile("book.txt")
	check(err)
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
