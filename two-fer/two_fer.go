package twofer

import "fmt"

// ShareWith returns name in a particular format
func ShareWith(name string) string {
	const ofm = ", one for me."
	const of = "One for "
	const ofy = "One for you"
	if len(name) > 0 {
		result := fmt.Sprintf("%s%s%s", of, name, ofm)
		return result
	}
	result := fmt.Sprintf("%s%s", ofy, ofm)
	return result
}
