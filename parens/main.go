package main

import (
	"fmt"
)

func main() {
	fmt.Println("All possible parentheses")

	var list []string

	addParen(&list, 4, 0, "")

	for _, paren := range list {
		fmt.Println(paren)
	}
}

func addParen(list *[]string, count int, rightRem int, parens string) {
	if count == 0 {
		for rightRem > 0 {
			rightRem--
			parens = parens + ")"
		}

		*list = append(*list, parens)
	} else {
		parens1 := parens + "("
		addParen(list, count-1, rightRem+1, parens1)

		if rightRem > 0 {
			parens2 := parens + ")"
			addParen(list, count, rightRem-1, parens2)
		}
	}
}
