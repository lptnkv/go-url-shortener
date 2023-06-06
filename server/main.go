package main

import (
	"fmt"
	"strings"
)

var counter int

func idToShortUrl(id int) string {
	alphabet := "abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	base := len(alphabet)
	var res strings.Builder
	fmt.Println(base)
	var digits []int
	for id > 0 {
		remainder := id % base
		digits = append(digits, remainder)
		id = id / base
	}
	for i := len(digits) - 1; i >= 0; i-- {
		char := alphabet[digits[i]]
		res.WriteByte(char)
	}
	return res.String()
}

type Record struct {
	originalUrl string
	shortUrl    string
}

func main() {

}
