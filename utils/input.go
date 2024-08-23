package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetStyle() [][]string {
	var chrs [][]string
	var currRow []string
	standard, err := os.ReadFile("styles/standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	Checkhash(string(standard))
	standard = standard[1:]
	tempChrs := strings.Split(string(standard), "\n\n")
	for j := 0; j < len(tempChrs); j++ {
		currRow = strings.Split(tempChrs[j], "\n")
		chrs = append(chrs, currRow)
	}
	return chrs
}
