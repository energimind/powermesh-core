package mongoquery

// singular returns the singular form of a plural word.
func singular(plural string) string {
	if plural == "" || plural[len(plural)-1] != 's' {
		return plural
	}

	if plural[len(plural)-3:] == "ies" {
		return plural[:len(plural)-3] + "y"
	}

	if plural[len(plural)-2:] == "es" {
		return plural[:len(plural)-2]
	}

	return plural[:len(plural)-1]
}
