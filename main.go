package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		c, err := Convert(os.Args[1])
		if err != nil {
			fmt.Println("invalid format")
			os.Exit(1)
		}

		fmt.Println(c)
		os.Exit(0)
	}

	fmt.Println("Entering interactive mode.")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c, err := Convert(scanner.Text())
		if err != nil {
			fmt.Println("invalid format")
			continue
		}

		fmt.Println(c)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Convert converts bidirectionally between base 10 and base 16 numbers.
func Convert(s string) (string, error) {
	if IsHex(s) {
		return HexToDec(s)
	} else {
		return DecToHex(s)
	}
}
