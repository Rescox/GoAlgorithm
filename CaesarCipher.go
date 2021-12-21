package main

import (
	"fmt"
	"unicode"
)

func caesarRotation(sText string, iShift int) string {
	sCipherText := []rune(sText)
	for i := 0; i < len(sText); i++ {
		if unicode.IsLower(rune(sText[i])) {
			if int(sText[i])+iShift > 122 {
				sCipherText[i] = rune(97 + (122+int(sText[i])+iShift)%122)
			} else if int(sText[i])+iShift < 97 {
				sCipherText[i] = rune(122 + (((int(sText[i])) - 97) + iShift))
			} else {
				sCipherText[i] = rune(int(sText[i]) + iShift)
			}
		} else if unicode.IsUpper(rune(sText[i])) {
			if int(sText[i])+iShift > 90 {
				sCipherText[i] = rune(65 + (90+int(sText[i])+iShift)%90)
			} else if int(sText[i])+iShift < 65 {
				sCipherText[i] = rune(90 + (((int(sText[i])) - 65) + iShift))
			} else {
				sCipherText[i] = rune(int(sText[i]) + iShift)
			}
		}
	}
	fmt.Println(sCipherText)
	return string(sCipherText)
}

func encodeCaesar(sPlainText string, iShift int) string {
	return caesarRotation(sPlainText, iShift)
}

func decodeCaesar(sCipherText string, iShift int) string {
	return caesarRotation(sCipherText, iShift*-1)
}

func main() {
	var sPlainText string
	var iShift int

	fmt.Println("Introduce your plain text")
	fmt.Scanln(&sPlainText)
	fmt.Println("Introduce the shift")
	fmt.Scanln(&iShift)
	fmt.Print(decodeCaesar(encodeCaesar(sPlainText, iShift), iShift))
}
