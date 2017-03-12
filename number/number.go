package number

import "fmt"

// Add function add together the numbers given and returns the result as a
// string
func Add(a, b int) string {
	num := a + b

	return intToString(num)
}

// Minus function take the value of b from the value of a and return the result
// as a string
func Minus(a, b int) string {
	num := a - b

	return intToString(num)
}

func intToString(num int) string {
	switch num {
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	default:
		return fmt.Sprintf("Cannot convert %v to a string", num)
	}
}
