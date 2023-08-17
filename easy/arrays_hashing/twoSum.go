package main


import "fmt"

func main() {
    nums := []int{2,7,11,15}
    
    result := twoSum(nums, 9)
    fmt.Println(result)
}


func twoSum(n []int, t int) []int {
    
    result := make([]int, 2)

    l := 0
    r := len(n) - 1
    
    for l < r {
        sum := n[l] + n[r]

        if sum == t {

            result[0] = l
            result[1] = r
            return result
        } else if sum < t {
            l++
        } else {
            r--
        }
    }

    return []int{-1, -1}
}
