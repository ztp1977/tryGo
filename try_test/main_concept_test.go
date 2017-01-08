package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// 必要な用語
//-bench regexp
//-benchmem
//-benchtime t
//-blockprofile block.out
//-blockprofilerate n
//-cover
//-covermode set,count,atomic
//-coverpkg pkg1,pkg2,pkg3
//-coverprofile cover.out
//-cpu 1,2,4
//-cpuprofile cpu.out
//-memprofile mem.out
//-memprofilerate n
//-outputdir directory
//-parallel n
//-run regexp
//-short
//-timeout t
//-v

// concept
// なぜエラーになったかを意識して、ケースを作成する
//

// go test
// go test -h

// functions
// t.Fail # エラーになっても続く
// t.FatalNow # そこでテストが止まる
// t.Skip # そのテストはやらない
// t.Log # 行数まで出力される
// go test -run Addr # Addrを含むテストのみ実行される
//

func TestSetsRemoteAddr(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", r.RemoteAddr)
	}))

	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("ReadAll err: %v", err)
	}
	ip := string(body)
	if !strings.HasPrefix(ip, "127.0.0.1:") && !strings.HasPrefix(ip, "[::1]:") {
		t.Fatalf("Excepted local addr; got %q", ip)
	}
}

func TestErrorMessage(t *testing.T) {
	if false {
		t.Error("show error message")
		t.Errorf("%s", "show error message by format")
	}

	t.Logf("%s", "write log by format")
}

func TestCart_Add(t *testing.T) {
	c := New()
	c.Add("りんご")
	c.Add("みかん")

	products := c.GetAll()
	if len(products) != 2 {
		t.Fatalf("商品の数が想定と異なる.商品数(%d)", len(products))
	}

	if products[0] != "りんご" {
		t.Errorf("りんごがカートにはいっていない。")
		t.Log("カートの中身: ", products)
	}

	if products[1] != "みかん" {
		t.Errorf("みかんがカートにはいっていない。")
		t.Log("カートの中身: ", products)
	}

	failMessage := EqualLength(products, 2)
	if failMessage != "" {
		t.Errorf(failMessage)
	}
}

func EqualLength(values ...interface{}) (failureMessage string) {
	actualValue := reflect.ValueOf(values[0])

	if actualValue.Kind() != reflect.Slice && actualValue.Kind() != reflect.Array {
		failureMessage = fmt.Sprintf("`%v` is not slice or array", values[0])
	}

	if actualValue.Len() != values[1] {
		failureMessage = fmt.Sprintf("Excepted `%v` length `%v` to equal `%v`", values[0], actualValue.Len(), values[1])
	}
	return failureMessage
}
