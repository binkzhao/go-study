package main

import "fmt"

func main() {
Work:
	for i := 0; i < 10; i++ {
		Work2:
		for j := 0; j < 10; j ++ {
			if i == 3 && j == 3 {
				break Work2
			}

			if i == 5 && j == 5 {
				//break
				break Work
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("done")
}
