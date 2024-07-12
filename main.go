package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]
	result := fileSelect(path)
	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, errs := file.WriteString(result)
	if errs != nil {
		fmt.Println("Failed to write to file:", errs)
		return
	}
}
