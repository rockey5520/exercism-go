// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

//Package proverb will print proverbs
package proverb

//Proverb will rhyme
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return []string{}
	}

	var result []string
	for i, val := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, "And all for the want of a "+string(rhyme[0])+".")
		} else {
			result = append(result, "For want of a "+val+" the "+string(rhyme[i+1])+" was lost.")
		}
	}
	return result
}
