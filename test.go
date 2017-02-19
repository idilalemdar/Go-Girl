package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	dat,err := ioutil.ReadFile("book.txt")
	check(err)
	var parsed []string = strings.Split(string(dat), "\n\n")
	var count int = 0

	for i := 0; i <= len(parsed)-1; i++ {
	    	if !strings.Contains(parsed[i], "Chapter") && (len(parsed[i]) != 0){
			fmt.Println(parsed[i])
			count++ 		
	}
}
	fmt.Println(count)
	fmt.Println(len(parsed))
}


//Counts the # of paragraphs correctly. (2063)