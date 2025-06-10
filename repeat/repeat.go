package iteration

import (
	"fmt"
	"strings"
)

const repeatCount = 5 // Repeat takes a string and returns it repeated 5 times

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// Repeat takes a character string as input and returns a new string where the input is repeated 5 times.
// Example:
//
//	Repeat("a") // returns "aaaaa"
func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
