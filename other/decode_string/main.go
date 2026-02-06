package decodestring

import (
	"strconv"
	"strings"
)

func DecodeString(input string) string {
	number := make([]int, 0)
	template := make([]string, 0)
	var inCnt int

	res := ""

	var tmpStr string
	for _, elem := range input {
		if elem >= '0' && elem <= '9' { // TODO: как обработать 0
			num, _ := strconv.Atoi(string(elem))
			number = append(number, num)

		} else if elem == '[' {
			//в стек кинуть строку
			template = append(template, tmpStr)
			tmpStr = ""
			inCnt++

		} else if elem >= 'a' && elem <= 'z' {
			tmpStr = tmpStr + string(elem)

		} else if elem == ']' {
			mul := number[len(number)-1]
			number = number[:len(number)-1]

			inCnt--
			if inCnt > 0 {
				// multp actual tmp string
				mlpStr := MulStr(mul, tmpStr)

				// glue prev left str
				prevStr := template[len(template)-1]
				template = template[:len(template)-1]
				tmpStr = prevStr + mlpStr
			} else {
				res = res + MulStr(mul, tmpStr)
				tmpStr = ""
			}
		}

	}

	if tmpStr != "" {
		res = res + tmpStr
	}
	return res
}

func MulStr(multiplier int, input string) string {
	res := make([]string, multiplier)

	for i := range multiplier {
		res[i] = input
	}
	return strings.Join(res, "")
}
