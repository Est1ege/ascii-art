package main

import (
	"ascii-art/utils"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	inputText := os.Args[1]
	//проверка входных данных есть ли они в аски
	cor := utils.CheckInput(inputText)
	if cor == true {
	//получение массива строк разделяя каждую строку 
	words := strings.Split(string(inputText), "\\n")
	//получение файла c стилями
	styleChrs := utils.GetStyle()
	//вывод
	utils.Output(styleChrs, words)
	}
}
