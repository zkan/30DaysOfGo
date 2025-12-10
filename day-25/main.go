package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// contents, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	// fmt.Println("Contents of file:", string(contents))

	f, err := os.Open("test.txt")

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
