package util

import "fmt"

// Check error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
