package main

import (
	"fmt"
	"strconv"
	"strings"
)

// convert converts bidirectionally between base 10 and base 16 numbers.
func Convert(s string) (string, error) {
	if IsHex(s) {
		return HexToDec(s)
	} else {
		return DecToHex(s)
	}
}

func DecToHex(dec string) (string, error) {
	v, err := strconv.ParseInt(dec, 10, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%#x", v), nil
}

func HexToDec(hex string) (string, error) {
	h := strings.TrimPrefix(hex, "0x")

	v, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", v), nil
}

func IsHex(s string) bool {
	return strings.HasPrefix(s, "0x")
}
