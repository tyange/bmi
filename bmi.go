package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const separator = "--------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightPrompt = "Please enter your height (m): "

func main() {
	// 1. Output information
	fmt.Println(mainTitle)
	fmt.Println(separator)

	// 2. Prompt for user input (weight + height)

	// CLI를 통해 몸무게와 키를 입력 받는다.
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// 3. Save that user input in variables

	// 입력 받은 몸무게와 키에 '\n'이 포함되어 있으므로 해당 문자열을 제거해준다
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	// 필요하지 않은 문자열은 제거되었지만, 여전히 weightInput, heightInput의 타입이 숫자(float)가 아니므로,
	// 해당 문자열을 숫자로 변환해준다. float64 fucntion을 사용할 수도 있다. 그러나 float64는
	// 숫자가 아닌 문자열을 변환하지 못한다. 여기서 strconv 모듈을 쓴다. strconv 모듈은
	// 모든 문자열(입력값)을 원하는 숫자 타입의 데이터로 반환하며, 해당 문자열이 숫자가 아닌 경우 error를 반환한다.
	// (그래서 반환값이 두 개임)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
