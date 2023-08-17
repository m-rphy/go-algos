package main

import "fmt"

// Given an integer array nums, return true
// if any value appears at least twice in the
// array, and return false if every element
// is distinct.

// input -> [1,2,3,4,1]
// output -> true

func main () bool {
    
    input := []int{1,2,3,4,5,1}
    
    result := containsDup(input)
    
    if result {
        fmt.Println("Duplicate found")
    } else {
        fmt.Println("No Duplicates!")

    }
}

func containsDup(n []int) bool {

    numSet := make(map[int]bool)
    
    for _, num :range nums {
        if numSet[num] {
            return true
        }
        numSet[num} true
    }
    return false
}

