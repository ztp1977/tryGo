// 任意の構造を[]interface{}に変換する
package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/k0kubun/pp"
)

type MyStruct struct {
	Id  int    `json:"id"`
	Str string `json:"str"`
}

// 1. Reflection goes from interface value to reflection object.
// 2. Reflection goes from reflection object to interface value.
// 3. To modify a reflection object, the value must be settable.

func doStruct() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func getStructTag() {
	it := new(MyStruct)

	rt := reflect.TypeOf(*it)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		json := field.Tag.Get("json")
		pp.Println(json)
	}
}

func jsonToMap() {
	it := &MyStruct{Id: 1, Str: "12132"}
	j, err := json.Marshal(it)
	if err != nil {
		pp.Println(err)
		return
	}
	pp.Println(j)

	var mp map[string]interface{}
	if err = json.Unmarshal([]byte(j), &mp); err != nil {
		panic(err)
	}

	pp.Println(mp)

}

func main() {

	doStruct()

	getStructTag()

	jsonToMap()

	fmt.Println("vim-go")
}
