package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	emailHash := make(map[string]string)

	for _, email := range emails {
		items := strings.Split(email, "@")
		localName := strings.SplitN(strings.Replace(items[0], ".", "", -1), "+", 2)[0]
		emailHash[fmt.Sprintf("%s@%s", localName, items[1])] = ""
	}

	return len(emailHash)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	n := numUniqueEmails(emails)
	fmt.Println(n)
}
