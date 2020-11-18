// map to json

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	s := []map[string]interface{}{}

	m1 := map[string]interface{}{"name": "John", "age": 10}
	m2 := map[string]interface{}{"name": "Alex", "age": 12}

	s = append(s, m1, m2)
	s = append(s, m2)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	str := string(b)
	fmt.Println(reflect.TypeOf(s), s)
	//fmt.Println(reflect.TypeOf(b), b)
	fmt.Println(reflect.TypeOf(str), str)

	//m := []map[string]interface{}{}
	m := make([]map[string]interface{}, 0)
	//必须为指针
	json.Unmarshal(b, &m)
	fmt.Println(reflect.TypeOf(m), m)
	fmt.Println("------------------")

	n := make(map[string]interface{})
	n["1"] = 111
	n["2"] = "222"

	fmt.Println(reflect.TypeOf(n), n)
	b, err = json.Marshal(n)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	z := make(map[string]interface{})
	json.Unmarshal(b, &z)
	fmt.Println(reflect.TypeOf(z), z)

}
