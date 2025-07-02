package util

import "fmt"

func Confirm(msg string) bool {
	fmt.Printf("%s [y/N]: ", msg)
	var s string
	_, _ = fmt.Scanln(&s)
	return s == "y"
}
