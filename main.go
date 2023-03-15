package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arabOrRim(err error, err2 error) bool {

	if err == nil && err2 == nil {
		return true
	}
	if err != nil && err2 != nil {
		return false
	} else {
		err := errors.New("wrong input")
		fmt.Println(err)
		os.Exit(0)
	}
	return false
}

func arab(a int, b int, sign string) int {
	switch sign {

	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	default:
		err := errors.New("wrong input")
		fmt.Println(err)
		os.Exit(0)
		return -1
	}
}

func romanToInt(roman string) int {
	var sum int
	var Roman = map[byte]int{'I': 1, 'V': 5, 'X': 10}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= Roman[byte(v)]
		} else {
			sum += Roman[byte(v)]
		}
	}
	return sum
}

func intToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func main() {
	var result int
	exp, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	arr := strings.Split(strings.TrimSpace(exp), " ")

	if len(arr) > 3 {
		err := errors.New("wrong input")
		fmt.Println(err)
		os.Exit(0)
	}

	num1, err := strconv.Atoi(arr[0])
	num2, err2 := strconv.Atoi(arr[2])

	sign := arr[1]

	arabCon := arabOrRim(err, err2)

	if arabCon { // проверка введены арабские или римские цифры
		if num1 > 10 || num2 > 10 {
			err := errors.New("numbers must be less or equal 10")
			fmt.Println(err)
			os.Exit(0)
		}
		result = arab(num1, num2, sign)
		fmt.Println(result)
	} else {
		num1, num2 = romanToInt(arr[0]), romanToInt(arr[2])
		if num2 >= num1 && sign == "-" {
			err := errors.New("There are no negative numbers in roman system (or 0) ")
			fmt.Println(err)
			os.Exit(0)
		}
		result = arab(num1, num2, sign)
		fmt.Println(intToRoman(result))
	}

}
