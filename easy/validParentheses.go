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

/*
func main () {
	//fmt.Print(os.Args[1:])
	for args := range os.Args {
		fmt.Println(":", args)
	}

	return
}
*/

func main () {
	if 2 > len(os.Args) {
		fmt.Println(true)
		return
	}

	var i, j int = 0, 0
	for i = 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}


	var stack [10] string

	var bracketsmap map[string]string
	bracketsmap = make(map[string]string)
	bracketsmap["{"] = "}"
	bracketsmap["["] = "]"
	bracketsmap["("] = ")"

	var match bool = true
	for i = 1; i < len(os.Args); i++ {
		b, ok := bracketsmap[os.Args[i]]
		if ok {
			stack[j] = b
			j++
		} else {
			if stack[j - 1] == os.Args[i] {
				stack[j - 1] = ""
				j--
			} else {
				match = false
				break
			}
		}
	}

	if 0 != j {
		match = false
	}

	fmt.Println(match)

	return
}
