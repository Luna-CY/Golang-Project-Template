package request

import (
	"fmt"
	"testing"
)

func Test_trimStringSpace(t *testing.T) {
	var s = " abc "
	fmt.Println(s)
	if err := trimStringSpace(&s); nil != err {
		panic(err)
	}
	fmt.Println(s)

	var m = map[string]any{
		"a": " abc ",
		"b": "def ",
		"c": []string{" ddd "},
		"d": map[string]string{"jkl": " mno "},
	}
	fmt.Println(m)
	if err := trimStringSpace(&m); nil != err {
		panic(err)
	}
	fmt.Println(m)

	var sli = []string{"abc ", "def ", "ghi "}
	fmt.Println(sli)
	if err := trimStringSpace(&sli); nil != err {
		panic(err)
	}
	fmt.Println(sli)

	var stru = struct {
		Abc string
	}{
		Abc: " abc ",
	}
	fmt.Println(stru)
	if err := trimStringSpace(&stru); nil != err {
		panic(err)
	}
	fmt.Println(stru)
}

func TestBindTrimSliceEmptyValueHandler(t *testing.T) {
	var ss = []string{"abc", "", "def", "ghi", ""}
	fmt.Println(ss)
	if err := BindHandlerTrimSliceEmptyValue(&ss); nil != err {
		panic(err)
	}
	fmt.Println(ss)

	var st = struct {
		Abc []string
	}{
		Abc: []string{"abc", "", "def", "ghi", ""},
	}
	fmt.Println(st)
	if err := BindHandlerTrimSliceEmptyValue(&st); nil != err {
		panic(err)
	}
	fmt.Println(st)
}
