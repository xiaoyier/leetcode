package array

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	num := []int{2, 8, 3, 4, 1, 0, 3, 6}
	target := 11

	result := twoSum(num, target)
	t.Log(result)
	t.Logf("%v", result)

	a := -1 << 16
	b := -1 ^ a
	t.Log(a, b)
}

type Handle func(val int)
type Handles []Handle

func ff(v int) {
	fmt.Print(v)
}

type CC struct {
	Name string
}

func Parse(dst interface{}) {
	err := json.Unmarshal([]byte(`{"Name": "Sycamor"}`), &dst)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAppend(t *testing.T) {

	var sss []int
	sss = append(sss, 100)

	bbb := map[string]int{"age": 17, "height": 180}
	aaa := bbb
	delete(aaa, "age")
	height := aaa["height"]
	height -= 5
	t.Log(aaa, bbb)

	var cc CC
	Parse(cc)
	t.Log(cc)

	pp := make(map[string][]CC)
	pp["tidy"] = make([]CC, 0)
	tidy := pp["tidy"]
	tidy = append(tidy, CC{Name: "Download"})
	t.Log(pp, tidy)

	z := map[string]int{
		"1": 1,
		"2": 2,
	}

	zz := z["1"]
	zz = 100
	t.Log(zz)

	var x interface{} = 111
	b, _ := x.(string)
	t.Log(b)

	var h Handles = nil
	l := len(h)
	t.Log(l)

	a := []Handle{ff}
	h = append(h, a...)

	t.Log(h)
}
