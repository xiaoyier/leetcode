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