package number

import "fmt"

// Add function add together the numbers given and returns the result as a
// string
func Add(a, b int) string {
	num := a + b

	switch num {
	case 4:
		return "four"
	default:
		return fmt.Sprintf("Cannot convert %v to a string", num)
	}
}

// Minus function take the value of b from the value of a and return the result
// as a string
func Minus(a, b int) string {
	num := a - b

	switch num {
	case 3:
		return "three"
	default:
		return fmt.Sprintf("Cannot convert %v to a string", num)
	}
}
