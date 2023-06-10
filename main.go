package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var dictMap = make(map[string]struct{})
	var files = []string{"total.txt"}
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		r := bufio.NewReader(bytes.NewReader(content))
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}
			pieces := strings.Split(string(line), " ")
			if len(pieces) > 1 {
				fmt.Println(pieces[1])
				dictMap[pieces[1]] = struct{}{}
			}
		}
	}

	f, err := os.OpenFile("dict.txt", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for w := range dictMap {
		if len([]rune(w)) > 6 {
			continue
		}
		f.WriteString(w + "\n")
	}
	f.Sync()
}
