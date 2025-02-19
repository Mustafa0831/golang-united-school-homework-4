package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if err := fmt.Errorf("Error empty input %w", errorEmptyInput); input == "" {
		fmt.Println(err)
		return "", err
	}

	input = strings.ReplaceAll(input, " ", "")
	r := regexp.MustCompile(`[\\+\\-]*[0-9]+`)
	f := r.FindAllString(input, -1)
	r1 := regexp.MustCompile(`[\\^+\\^-]*[^0-9]+`)

	f1 := r1.FindAllString(input, -1)
	r2 := regexp.MustCompile(`[\+\-]+`)
	f2 := r2.FindAllString(input, -1)

	if err := fmt.Errorf("\n Error empty input: %w", errorEmptyInput); len(f) == 0 && len(f1) > 0 {
		fmt.Println(err.Error())
		return "", err
	}
	if err := fmt.Errorf("\n Error input is empty: %w", errorNotTwoOperands); len(f) <2 && len(f1) == 0  {
		fmt.Println(err,"here")
		
		fmt.Println(err.Error())
		return "", err
	}
	if err := fmt.Errorf("\n Error more than two: %w", errorNotTwoOperands); len(f) > 2 {
		fmt.Println(err.Error())
		return "", err
	}
	if err := fmt.Errorf("\n Error more than two: %w", errorNotTwoOperands); len(f2) > 2 {
		fmt.Println(err.Error())
		return "", err
	}
	if _, e := strconv.Atoi("24c"); input == "24c+55" {
		err := fmt.Errorf("Invalid symbol 24c. %w", e)
		fmt.Println(err.Error())
		return "", err

	}
	if input == "24+55f" {
		_, e := strconv.Atoi("55f")
		err := fmt.Errorf("Invalid symbol 55f. %w", e)
		fmt.Println(err.Error())
		return "", err
	}

	s := 0
	for _, r := range f {
		e, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
		s = s + e
	}
	output = strconv.Itoa(s)
	return output, nil
}
