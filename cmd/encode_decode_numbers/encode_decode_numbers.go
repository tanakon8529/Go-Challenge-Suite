package main

import (
	"fmt"
)

// decodeInput decodes the encoded string based on 'L', 'R', and '=' operations.
func decodeInput() []int {
	var encodedStr string
	if _, err := fmt.Scanln(&encodedStr); err != nil {
		return nil
	}

	decodedValues := make([]int, len(encodedStr)+1)
	runes := []rune(encodedStr)

	for i, char := range encodedStr {
		switch char {
		case 'L':
			for j := i; j >= 0; j-- {
				switch runes[j] {
				case 'L':
					adjustForL(decodedValues, j, j+1)
				case 'R':
					adjustForR(decodedValues, j, j+1)
				case '=':
					adjustForE(decodedValues, j, j+1)
				}
			}
		case 'R':
			adjustForR(decodedValues, i, i+1)
		case '=':
			adjustForE(decodedValues, i, i+1)
		}
	}
	return decodedValues
}

// adjustForL ensures the left value is greater than the right.
func adjustForL(values []int, left, right int) {
	if values[left] <= values[right] {
		values[left] = values[right] + 1
	}
}

// adjustForR ensures the right value is greater than the left.
func adjustForR(values []int, left, right int) {
	if values[right] <= values[left] {
		values[right] = values[left] + 1
	}
}

// adjustForE ensures the right value equals the left.
func adjustForE(values []int, left, right int) {
	values[right] = values[left]
}

func main() {
	fmt.Println(decodeInput())
}
