package main

import (
	"fmt"
	"unicode"
)

//import "fmt"

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

var r = []rune{}

func (c *NumericInput) Add(a rune) {
	res := unicode.IsDigit(a)
	if res {
		r = append(r, a)
	}
}
func (c *NumericInput) GetValue() string {
 return	 string(r)
	//return string(r[len(r)-1])
}
func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
