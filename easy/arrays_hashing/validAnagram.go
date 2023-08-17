package main

import "fmt"


func main() {

    str_1 := "listen"
    str_2 := "silent"

    result := areAnagrams(str_1, str_2)

    if result {
        fmt.Println("The strings are anagrams.")
    } else {
        fmt.Println("The strings are not anagrams.")
    }
}

// Helper function
func areAnagrams(s1 string, s2 string) bool {
    
    if len(s1) != len(s2) {
        return false
    }

    // accumulators for ASCII codes
    acc_1 := 0
    acc_2 := 0

    // Add the ASCII codes of each string
    for _, char := range s1 {
        acc_1 += int(char)
    }

    for _, char := range s2 {
        acc_2 += int(char)
    }

    // return the boolean of their equivalence
    return acc_1 == acc_2

}
