func numUniqueEmails(emails []string) int {
	emailHash := make(map[string]string)

	for _, email := range emails {
		items := strings.Split(email, "@")
		localName := strings.SplitN(strings.Replace(items[0], ".", "", -1), "+", 2)[0]
		emailHash[fmt.Sprintf("%s@%s", localName, items[1])] = ""
	}

	return len(emailHash)
}