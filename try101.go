package main

import (
		"fmt"
		"strings"
		"regexp"
)

func WordCount(s string) map[string]int {

	re := regexp.MustCompile(` +`)
	words := re.Split(s, -1)
	wordCount := make(map[string]int)
	for i:=0 ; i < len(words) ; i++ { 
		wordNoComma := strings.Trim(words[i],",")
		wordCount[wordNoComma]++
    }
        
    return wordCount
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quack like a duck, then it probably is a duck"
	w := WordCount(s)
	fmt.Println(w)
}
