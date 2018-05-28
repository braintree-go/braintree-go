package testhelpers

// Contains checks whether a slice of strings contains
// a given string
func Contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
