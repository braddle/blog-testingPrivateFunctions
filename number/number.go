package number

import "fmt"

// Add function add together the numbers given and returns the result as a string
func Add(a, b int) string {
	num := a + b

	switch num {
	case 4:
		return "four"
	default:
		return fmt.Sprintf("Cannot convert %v to a string", num)
	}
}
