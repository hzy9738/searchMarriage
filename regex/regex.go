package main

import (
	"regexp"
	"fmt"
)

const  text = `
my email is hzy9738@126.com
hhh@163.com
zzz@gmail.com
yyy@ooowin.com
`

func main() {
	re := regexp.MustCompile(
		`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(text,-1)
	fmt.Println(match)
}
