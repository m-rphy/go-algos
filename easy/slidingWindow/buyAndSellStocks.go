package main

import (
	"fmt"
	"math"
)

// You are given an array prices where
// prices[i] is the price of a given
// stock on the ith day.

// You want to maximize your profit by
// choosing a single day to buy one
// stock and choosing a different day
// in the future to sell that stoc

func main() {

    nums := []int {1,2,4,3,8,5,4,9}

    result := maxProfit(nums)

    fmt.Printf("The max profit is: %v\n", result)
}

func maxProfit(p []int) int {

    profit := 0

    lowest := p[0]

    for _, price := range p {
        if price < lowest {
            lowest = price
        }
        profit = max(profit, price - lowest)
    }

    return profit
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

