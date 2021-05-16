package stack

import (
	"fmt"
	"testing"
)


func TestValidKuohao(t *testing.T) {

	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("[](()"))
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("{[("))
	fmt.Println(isValid("}])("))

}

func TestEvalRPN(t *testing.T) {

	str := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(str))

}

func TestBasicCalculator(t *testing.T) {

	//s := "1 + 1 - (2-9+3)"
	//fmt.Println(calculate(s))   //6
	fmt.Println(calculate("22435565"))
}
