package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("instance.sat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expression [][]int

	var clauseIndex = 0

	for scanner.Scan() {
		line := scanner.Text()
		switch firstChar := line[:1]; firstChar {
		case "c":
			// ignore comment lines
		case "p":
			// read numVars and numClauses from p(repare?) line
			tokens := strings.Split(line, " ")
			totalNumVars, _ := strconv.Atoi(tokens[2])
			totalNumClauses, _ := strconv.Atoi(tokens[3])
			fmt.Println("totalNumVars:", totalNumVars)
			fmt.Println("totalNumClauses:", totalNumClauses)
			expression = make([][]int, totalNumClauses)
		default:
			// read variables into a list of lists, representing an AND of OR clauses
			tokens := strings.Split(line, " ")
			numVars := len(tokens)-1
			expression[clauseIndex] = make([]int, numVars)
			for varIndex := 0; varIndex < numVars; varIndex++ {
				intRepr, _ := strconv.Atoi(tokens[varIndex])
				expression[clauseIndex][varIndex] = intRepr
			}
			clauseIndex++
		}
	}

	fmt.Println(expression)
}
