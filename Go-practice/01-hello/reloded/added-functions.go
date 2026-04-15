package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func convertHexBin(text string) string {

	word := strings.Fields(text)

	
	result := []string{}

	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(word[i-1], 16, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprintf("%d", num)
			}
		} else if word[i] == "(bin)" && i > 0 {
			num, err := strconv.ParseInt(word[i-1], 2, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprintf("%d", num)
			}
		}else {
			result = append(result, word[i])
		}
	}
	return strings.Join(result, " ")
}

func caseUp(text string) string{

	word := strings.Fields(text)

	result := []string{}

	for i := 0; i < len(word); i++{
		if word[i] == "(up)" && i > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)])
		}else{
			result = append(result,  word[i])
		}
	}
	return strings.Join(result, " ")
}

func caseLow(text string) string{

	word := strings.Fields(text)

	result := []string{}

	for i := 0; i < len(word); i++{
		if word[i] == "(low)" && i > 0{
			result[len(result)-1] = strings.ToLower(result[len(result)])
		}else{
			result = append(result, word[i])
		}

	}
	return strings.Join(result, " ")
}