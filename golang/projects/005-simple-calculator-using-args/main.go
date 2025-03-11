package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	calculator()
}

func calculator() {
	args := os.Args

	numA, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("error!\n%v is not integer.\n", numA)
		os.Exit(2)
	}
	
	numB, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("error!\n%v is not integer.\n", numB)
		os.Exit(2)
	}

	result := numA + numB;
	fmt.Printf("%d + %d = %d\n", numA, numB, result)
}