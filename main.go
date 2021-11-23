package main

import "os"
import "fmt"
import "strings"

func setDictionary() map[string]string {
	character := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}

	return character
}

func isAlphaNumeric(input string) bool {
	alphaNumeric := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range alphaNumeric {
		if string(char) == input{
			return true
		}

	}

	return false
}

func convert(dictionary map[string]string, input string) string{
	var stringToConvert string = ""
	for _, char := range input {
		lowerCaseChar := strings.ToLower(string(char))
		if !isAlphaNumeric(lowerCaseChar){ 
			stringToConvert += string(char)
		 }
		stringToConvert += dictionary[lowerCaseChar] + " "
	}

	return stringToConvert
} 

func main(){
	input := os.Args[1]
	output := convert(setDictionary(), input) 

	fmt.Println(output)
}