/*
# Given a string containing just the characters '(',')','{','}','[',']', determin if the input string is valid.

#An input string is valid if:

#Open brackets must be closed by the same of brackets
#open brackets must be closed by the correct order.
#Note that an empty string is also considerd valid

#Example 1:
#Input: "()"
#Output: true

#Example 2:
#Input: "()[]{}"
#Output: true

#Example 3:
#Input: "(]"
#Output: false

#Example 4:
#Input: "([)]"
#Output: failse

#Example 5:
#Input: "{[]}"
#Output: true
*/

/*
package main

import "fmt"

func main(){
	fmt.Println("Hello, Golang")
}
*/

package main

import (
	"fmt"
	"os"
)

func main () {
	//fmt.Print(os.Args[1:])
	for args := range os.Args {
		fmt.Println(":", args)
	}

	return
}
