package main

import (
	"fmt"
	"regexp"
)

func main() {

	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)
	var method1 = regex.FindString(text)
	var method2 = regex.FindAllString(text, 2)
	fmt.Println(method1, method2)
}
