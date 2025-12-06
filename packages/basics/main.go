// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World, from Gopher!")
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	fmt.Println(
		"The date and time now, in Tokyo, is:",
		time.Now().In(tz).Format(time.RFC1123),
	)
}
