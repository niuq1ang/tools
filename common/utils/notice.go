package utils

import (
	"fmt"
)

func Notice(str string) {
	fmt.Println("===========================")
	defer fmt.Println("===========================")
	fmt.Println(str)
}
