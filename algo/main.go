package main

import "strings"

func main() {
	strSlices := []string{"h", "e", "l", "l", "o"}

	var strb strings.Builder
	for _, str := range strSlices {
		strb.WriteString(str)
	}
	print(strb.String())
}