// struct:cols -> []interface{}
package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx/reflectx"
	"github.com/k0kubun/pp"
)

type MyStruct struct {
	Id   int
	Name string
	ts   time.Time
}

func InterfaceDest(dest interface{}) (ret []interface{}) {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		panic("InterfaceSlice() given a non-slice type")
	}

	v = reflect.Indirect(v)
	return ret
}

type Rows struct {
	*sql.Rows
	unsafe bool
	Mapper *reflectx.Mapper
	// these fields cache memory use for a rows during iteration w/ structScan
	started bool
	fields  [][]int
	values  []interface{}
}

func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		pp.Println(v.Kind().String())
		pp.Println("12131123")
		return errors.New("argument not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := reflectx.FieldByIndexes(v, traversal)
		if ptrs {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
	return nil
}

func main() {

	v := &MyStruct{Id: 100, Name: "name", ts: time.Now()}
	pp.Println(v)
	json, _ := json.Marshal(v)
	pp.Println(string(json))
	pp.Println("1231312")
	r := new(Rows)
	rv := reflect.ValueOf(v)
	rvv := fieldsByTraversal(rv, r.fields, r.values, true)
	pp.Println(v)
	pp.Println(rv)
	pp.Println(rvv)
	pp.Println(r.fields)
	pp.Println(r.values)

	fmt.Println("vim-go")
}
