package main

import "fmt"

func uniqueNames(a, b []string) []string {
	var result []string
	for _, v := range a {
		f := contains(result, v)
		if !f {
			result = append(result, v)
		}
	}

	for _, v := range b {
		f := contains(result, v)
		if !f {
			result = append(result, v)
		}
	}
	return result
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
