package main

import (
	//"context"

	"github.com/rusfort/abstract-go-func/functions"
)

func main() {
	println("hello")
	//ctx := context.Background()
	f1 := functions.AbstractFunc(func(params ...interface{}) (interface{}, error) {
		println("called f1")
		return nil, nil
	})

	m := make(functions.FuncMap)
	m["f1"] = f1

	f, ok := m["f1"]
	if ok {
		f.(functions.AbstractFunc)()
	}
}
