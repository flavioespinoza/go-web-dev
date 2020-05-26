package main

import (
	"fmt"
)

func main() {

	sages := map[string]string{
		"American": "Cheese Burger and Fries",
		"Indian":   "Tiki Masala",
		"Japanese": "Yaki Soba",
		"Mexican":  "Chile Verde",
		"Russian":  "Vodka",
	}

	for k, v := range sages {
		fmt.Println(k, " - ", v)
	}

}
