package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("instance.sat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		switch firstChar := line[:1]; firstChar {
		case "c":
			// ignore comment lines
		case "p":
			tokens := strings.Split(line, " ")
			numVars := tokens[2]
			numClauses := tokens[3]
			fmt.Println("numVars:", numVars)
			fmt.Println("numClauses:", numClauses)
		default:
			tokens := strings.Split(line, " ")
			for i := 0; i < len(tokens)-1; i++ {
				fmt.Println(tokens[i])
			}
		}
	}
}
