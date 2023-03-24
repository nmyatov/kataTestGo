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

func romanToInt(s string) int {
	var RomanNumerals = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}

func intToRoman(num int) string {
	var roman string
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, conversion := range conversions {
		for num >= conversion.value {
			roman += conversion.digit
			num -= conversion.value
		}
	}

	return roman
}

func main() {
	allowed_romans := [...]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
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

	arabCon := arabOrRim(err, err2)

	sign := arr[1]

	if arabCon { // проверка введены арабские или римские цифры
		if num1 > 10 || num2 > 10 {
			err := errors.New("numbers must be less or equal 10")
			fmt.Println(err)
			os.Exit(0)
		}
		if num1 < 1 || num2 < 1 {
			err := errors.New("numbers can't be less 1")
			fmt.Println(err)
			os.Exit(0)
		}
		result = arab(num1, num2, sign)
		fmt.Println(result)
	} else {
		var count1, count2 int

		for _, el := range allowed_romans {
			if el == arr[0] {
				count1++
			}
		}

		for _, el := range allowed_romans {
			if el == arr[2] {
				count2++
			}
		}

		if count1 == 0 || count2 == 0 {
			err := errors.New("wrong input")
			fmt.Println(err)
			os.Exit(0)
		}

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
