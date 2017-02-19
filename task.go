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
	for i := 1; i <= len(chapParsed) - 1; i++ {
		var paragParsed []string = strings.Split(chapParsed[i+1], "\n\n") 
		var index int = 0
		for j := 1; j <= len(paragParsed) - 1; j++ {
			if len(paragParsed[j]) != 0 {
				Book[i][index] = paragParsed[j]
				index++
			}			
		}
	}
}
func query(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("c") //Get Chapter from url as string. you have to convert it to int by using strconv.Atoi
	p := q.Get("p") //Get Paragraph from url as string, you have to convert it to int
	result := ""
	chap,_:= strconv.Atoi(c)
	para,err:= strconv.Atoi(p)
	if err != nil { //only chapter is requested
		result = strings.Join(Book[chap-1], "\n\n")
	} else {
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
