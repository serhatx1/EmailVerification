package main

import (
	"fmt"
	"regexp"
)

func main() {
	emails := []string{
		"serhat@gmail.com",
		"wrongonpurpose",
		"blabla@hotmail.com",
	}
	for _, email := range emails {
		if isEmailValid(email) {
			fmt.Println("%s:Email is valid", email)
		} else {
			fmt.Println("%s:Email is invalid", email)
		}
	}
}
func isEmailValid(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9._%+-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
