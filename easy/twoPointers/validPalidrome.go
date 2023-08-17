package main

import "fmt"


func main() {

    check := "racecar"

    result := valid(check)
    if result {
        fmt.Println("Valid palindrome")
    } else {
        fmt.Println("Not an palidrome")
    }
}


func valid(s string) bool {

    l := 0
    r := len(s) - 1

    for l < r {

        if s[l] != s[r] {
            return false
        }
        r--
        l++
    }
    return true
}
