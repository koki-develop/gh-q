package util

import "fmt"

func Confirm(msg string) bool {
	fmt.Printf("%s [y/N]: ", msg)
	var s string
	fmt.Scanln(&s)
	return s == "y"
}
