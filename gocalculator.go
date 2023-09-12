package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type calculator interface {
	Result(input string, result *float64) (float64, string)
	add(a float64, b float64) float64
	subtract(a, b float64) float64
	multiply(a, b float64) float64
	divide(a, b float64) float64
}

type mycalculator struct {
	name string
	id   int
}

func maincal() {
	defer fmt.Println("calculator exited!")

	var cal calculator
	cal = &mycalculator{
		name: "op",
		id:   1,
	}
	input := ""
	var result float64
	fmt.Scanln(&input)
	cal.Result(input, &result)
	fmt.Printf("%.8f\n", result)
	fmt.Println("按 Enter 键继续...")
	fmt.Scanln()

}

func (m *mycalculator) Result(input string, result *float64) (float64, string) {
	var s1, s2 float64
	var c string
	pattern := "[-+]?\\d+\\.\\d+|[-+]?\\d+"
	regex1, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("regexp compile error!") // 处理编译错误
	}
	match1 := regex1.FindString(input)
	s1, _ = strconv.ParseFloat(match1, 64)
	input = strings.Replace(input, match1, "", 1)

	pattern = "[-*+/]"
	regex2, err := regexp.Compile(pattern)
	c = regex2.FindString(input)
	input = strings.Replace(input, c, "", 1)

	pattern = "[-+]?\\d+\\.\\d+|[-+]?\\d+"

	regex1, err = regexp.Compile(pattern)
	if err != nil {
		fmt.Println("regexp compile error!") // 处理编译错误
	}
	match2 := regex1.FindString(input)
	s2, _ = strconv.ParseFloat(match2, 64)
	input = strings.Replace(input, match2, "", 1)
	//cut first float64+"+-*/"+float64 from input
	switch c {
	case "+":
		*result = m.add(s1, s2)
		break
	case "-":
		*result = m.subtract(s1, s2)
		break
	case "*":
		*result = m.multiply(s1, s2)
		break
	case "/":
		*result = m.divide(s1, s2)
		break
	} //result=float64+"+-*/"+float64

	if input != "" {
		tempstr := strconv.FormatFloat(*result, 'f', -1, 64)
		inputbuf := []byte(input)
		tempbuf := []byte(tempstr)
		input = string(append(tempbuf, inputbuf...))
		*result, input = m.Result(input, result)
	}
	return *result, input
}

func (m *mycalculator) add(a float64, b float64) float64 {
	return a + b
}

func (m *mycalculator) subtract(a float64, b float64) float64 {
	return a - b
}

func (m *mycalculator) multiply(a float64, b float64) float64 {
	return a * b
}

func (m *mycalculator) divide(a float64, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}

func Newmycalculator(name string, id int) *mycalculator {
	return &mycalculator{
		id:   id,
		name: name,
	}
}
