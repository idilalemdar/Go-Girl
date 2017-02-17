package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
)

var Book [][]string

func check(e error) {
    if e != nil {
        panic(e)
    }
} 

func readBook(filePath string) {
	dat,err := ioutil.ReadFile(filePath)
	check(err)
	var chapParsed []string = strings.Split(string(dat),"Chapter ") //parsing the text into its chapters
	for i := 0; i <= len(chapParsed) - 2; i++ {
		var paragParsedRaw []string = strings.Split(chapParsed[i+1], "\n\n") //parsing the text into its paragraphs, but there will be blank paragraphs as well
		var paragParsed []string
		var index int = 0
		for j := 1; j <= len(paragParsedRaw) - 1; j++ {
			if !strings.ContainsAny(paragParsedRaw[j], "* * * * *") && len(paragParsedRaw[j]) != 0 {
				paragParsed[index] = paragParsedRaw[j]
				index++
			}			
		}
		Book[i] = paragParsed
	}
}
func query(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("c") //Get Chapter from url as string. you have to convert it to int by using strconv.Atoi
	p := q.Get("p") //Get Paragraph from url as string, you have to convert it to int
	result := ""
	chap,_:= strconv.Atoi(c)
	if len(p) == 0 { //only chapter is requested
	/*	for i := 0; i <= len(Book[chap-1]-1; i++ {
			
		} */ //will be completed later
	} else {
		para,_:= strconv.Atoi(p)
		result = Book[chap-1][para-1] 	
	}
	fmt.Fprint(w, result)
}
func count(w http.ResponseWriter, r *http.Request) {
	chapCount := len(Book)
	paraCount := 0
	for i:= 0; i <= chapCount - 1; i++ {
		paraCount += len(Book[i])
	}
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
