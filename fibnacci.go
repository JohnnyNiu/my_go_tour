package main

import "fmt"

//(0, 1, 1, 2, 3, 5, ...).
func fibonacci() func() int {
	i := 0
	j := 0
	count := 0
	return func() int {
		result := 0
		if count == 0 {
			result = 0
			count = count + 1
		} else if count == 1 {
			result = 1
			i = 0
			j = 1
			count = count + 1
		} else {
			result = i + j
			i = j
			j = result
		}
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
