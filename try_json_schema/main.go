package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

func descrialize(data []byte, typ reflect.Type) (interface{}, error) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	payload := reflect.New(typ).Interface()
	if err := json.Unmarshal(data, payload); err != nil {
		return nil, err
	}

	return payload, nil
}

type Foo struct {
	Foo string
	Bar string
}

func main() {
	data := `{"foo": "bar", "bar2": "baz"}`
	it := map[string]string{}
	F := &Foo{}
	fmt.Println(descrialize([]byte(data), reflect.TypeOf(it)))
	fmt.Println(descrialize([]byte(data), reflect.TypeOf(F)))
	fmt.Println(it["foo"])
	fmt.Println(F.Foo)


	// jsonの使い方

	//
}
