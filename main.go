package main

import "fmt"

var stack [8000]byte

func main() {
	for j := 0; j < 8000; j++ {
		stack[j] = 0
	}
	var run string
	fmt.Println("Enter the code to execute")
	fmt.Scanf("%v", &run)

	var temp byte
	i := 0
	for j := 0; j < len(run); j++ {
		v := run[j]
		switch v {
		case 172:
			fmt.Print(stack[i])
		case 't':
			fmt.Print(stack[i])
			stack[i] = 0
		case '*':
			fmt.Scanf("%v", &temp)
			stack[i] = temp
		case 'o':
			stack[i]++
		case 'u':
			stack[i]--
		case 'e':
			i--
		case 'r':
			i++
		case 'a':
			if stack[i] != 0 {
				j = findEnd(run, j)
				continue
			}
			i++
		case 's':
			if stack[i] != 0 {
				j = findStart(run, j)
				continue
			}
			i++
		}
	}
	fmt.Println("\n\nThe pointer is at position", i, "in the stack")
	fmt.Printf("The stack looks like:\n%v", stack)
}

func findEnd(input string, position int) int {
	ret := position
	c := 1
	for c > 0 {
		ret++
		if input[ret] == 'a' {
			c++
		}
		if input[ret] == 's' {
			c--
		}
	}
	return ret - 1
}

func findStart(input string, position int) int {
	ret := position
	c := 1
	for c > 0 {
		ret--
		if input[ret] == 'a' {
			c--
		}
		if input[ret] == 's' {
			c++
		}
	}
	return ret
}
