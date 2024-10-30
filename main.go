package main

import (
	"fmt"
	"strings"
)

// ReverseString reverses the given string with underscores between characters, ignoring spaces.
func ReverseString(str string) string {
	var reversed []rune

	// Reverse the string and collect characters
	for i := len(str) - 1; i >= 0; i-- {
		char := rune(str[i])
		reversed = append(reversed, char)
	}

	var result []string
	// Prepare the result with underscores
	for i, char := range reversed {
		result = append(result, string(char))
		if char != ' ' && i < len(reversed)-1 && reversed[i+1] != ' ' {
			result = append(result, "_") // Add underscore between characters, not after spaces
		}
	}

	// Join the result into a single string
	return strings.Join(result, "")
}

// debugging
func main() {
	fmt.Println(ReverseString("Hello World"))     // output: "d_l_r_o_W o_l_l_e_H"
	fmt.Println(ReverseString("I am a student"))   // output: "t_n_e_d_u_t_s a_m_a_I"
}
