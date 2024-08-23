package utils

import "fmt"

func Output(chrs [][]string, words []string) {
	for i, currWord := range words {
		if i == 0 && currWord == "" && len(words) == 2 {
			continue
		}
		//проверка на /n
		if currWord == "" {
			fmt.Println()
		} else {
			// используем тройной цикл для вывода
			for i := 0; i < 8; i++ {
				for _, ch := range currWord {
					if ch >= 32 && ch <= 127 {
						for j := 0; j < len(chrs[ch-32][i]); j++ {
							fmt.Print(string(chrs[ch-32][i][j]))
						}
					}
				}
				fmt.Println()
			}
		}
	}
}
