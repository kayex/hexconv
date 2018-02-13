package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
