package validate

// implement the actor function
// it should return true if the id is valid
// and false otherwise
func Actor(id string) bool {
	return validate(id, "nm")
}

// implement the movie function
// it should return true if the id is valid
// and false otherwise
// use validate function
func Movie(id string) bool {
	return validate(id, "tt")
}

// extract common code to a function
// and use it in both functions
func validate(id string, prefix string) bool {
	if len(id) != 9 {
		return false
	}

	if id[0:2] != prefix {
		return false
	}

	for _, c := range id[2:] {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}
