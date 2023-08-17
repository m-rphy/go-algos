package main


func main () {

    result := validParens("{[()]}")
    
    if result {
        println("Valid parentheses")
    } else {
        println("Invalid parentheses")
    }
}


func validParens(s string) bool {
    stack := make([]rune, len(s))

    hash := map[rune]rune{
        "}": "{",
        "]": "[",
        ")": "(",
    }

    // range over the input string
    for _, char := range s {

        // if the key is in the hash
        if el, ok := hash[char]; ok {

            // if the stack is empty of the 
            // the last el on the stack doesn't match
            // the parathesis...
            if len(stack) == 0 || stack[len(stack)-1] != el {
                return false
            }

            // "pop" the rune from the stack
            stack = stack[:len(stack) - 1]

        } else {
            // "push" to the stack
            stack = append(stack, char)
        }
    }

    return len(stack) == 0
}
