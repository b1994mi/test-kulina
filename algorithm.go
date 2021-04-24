package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Ismail's algorithm answer!")
	fmt.Println("Menu:")
	fmt.Println("1. Sock Merchant Algorithm")
	fmt.Println("2. Counting Valleys")
	fmt.Println("3. Number Breakdown")
	fmt.Println("4. Andrew's Trip")
	fmt.Println("select a number:")
	var selection int
	fmt.Scanln(&selection)
	switch selection {
	case 1:
		sockMerchant()
	case 2:
		countingValleys()
	case 3:
		numberBreakdown()
	}
}

func sockMerchant() {
	fmt.Println("input n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("input ar: ")
	ar := make([]int, n)

	// Use scanner and loop, otherwise it won't read space-separated token.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Scan the text input and transform into slice of strings.
		inputSlice := strings.Fields(scanner.Text())

		// Simple handling if the number of space-separated input less than n.
		if len(inputSlice) > n || len(inputSlice) < n {
			println("invalid! input ar:")
			continue
		}

		// Convert every element of inputSlice to int.
		for i, num := range inputSlice {
			nu, _ := strconv.Atoi(num)
			ar[i] = nu
		}

		// Put conditional just to remove warning.
		if len(inputSlice) == n {
			break
		}
	}

	// Error handling.
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	// A slice to keep note which index inside ar that has been identified a pair.
	indexOfSameEl := make([]int, 0)

	for i := 0; i < len(ar); i++ {
		isElmTaken := false

		// Check whether this element has been tagged as a pair.
		for _, u := range indexOfSameEl {
			// If the current element has been tagged as a pair, skip this iteration.
			if i == u {
				isElmTaken = true
			}
		}

		if isElmTaken {
			continue
		}

		for j := i + 1; j < n; j++ {
			if ar[i] == ar[j] {
				indexOfSameEl = append(indexOfSameEl, i)
				indexOfSameEl = append(indexOfSameEl, j)
				break
			}
		}
	}

	// Length of slice indexOfSameEl is expected to be even.
	println(len(indexOfSameEl) / 2)
}

func countingValleys() {
	fmt.Println("input n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("input s: ")
	s := make([]string, n)

	// Use scanner and loop, otherwise it won't read space-separated token.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Scan the text input and transform into slice of strings.
		inputSlice := strings.Fields(scanner.Text())

		// Simple handling if the number of space-separated input less than n.
		if len(inputSlice) > n || len(inputSlice) < n {
			println("invalid! input s:")
			continue
		}

		// Check if input is string U or D.
		isValid := true
		for i, str := range inputSlice {
			if str != "U" && str != "D" {
				isValid = false
			}
			s[i] = str
		}

		// Handling if input is not U or D.
		if !isValid {
			println("invalid! input s:")
			continue
		}

		// Put conditional just to remove warning.
		if len(inputSlice) == n {
			break
		}
	}

	elevation, countValley, hasEnteredValley, hasExitedValley := 0, 0, false, false

	for i := 0; i < n; i++ {
		elevationBefore := elevation
		if s[i] == "D" {
			elevation--
		} else {
			elevation++
		}
		if elevationBefore == 0 && elevation < 0 {
			hasEnteredValley = true
		}
		if elevationBefore < 0 && elevation == 0 {
			hasExitedValley = true
		}
		if hasEnteredValley && hasExitedValley {
			countValley++
			hasEnteredValley, hasExitedValley = false, false
		}
	}

	println(countValley)
}

func numberBreakdown() {
	fmt.Println("input number: ")
	var input string
	fmt.Scanln(&input)
	input = strings.Replace(input, ".", "", -1)

	for i := len(input) - 1; i >= 0; i-- {
		fmt.Printf("%c", input[len(input)-i-1])
		for j := 0; j < i; j++ {
			print("0")
		}
		println()
	}
}
