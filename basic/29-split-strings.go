package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// split String using the split() function
	s1 := "This,is,a,delimited,string"
	v1 := strings.Split(s1, ",")
	fmt.Println(v1) // [This is a delimited string]

	// string Splitting using the SplitN function
	s2 := "This.is.a.delimited.string"
	v2 := strings.SplitN(s2, ".", 5)
	fmt.Println(v2) // [This is a delimited string]

	// String split using whitespace
	s3 := "This is a string containing whitespaces"
	v3 := strings.Fields(s3)
	fmt.Println(v3) // [This is a string containing whitespaces]

	// String Split without Removing Delimiters
	s41 := "String,with,delimiters."
	v41 := strings.SplitAfter(s41, ",") // It doesn't removes delimiters when splitting
	fmt.Println(v41)                    // [String, with, delimiters.]

	s42 := "averylargeword"
	v42 := strings.SplitAfter(s42, "") // generate a slice of individual characters from a string
	fmt.Println(v42)                   // [a v e r y l a r g e w o r d]

	// String split using the regexp
	s5 := "xyzxyzxyzxyzxyz"
	t5 := regexp.MustCompile(`[y]`) // backticks are used here to contain the expression
	v5 := t5.Split(s5, -1)          // second arg -1 means no limits for the number of substrings
	fmt.Println(v5)                 // [x zx zx zx zx z]
}
