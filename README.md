# abstract-go-func
Lib for abstract functions and methods in go

Try this out!

```go
package main

import (
	"github.com/rusfort/abstract-go-func/functions"
)

func main() {
	println("hello")

	// create your new abstract function
	f1 := functions.NewAbstractFunc(func(params ...interface{}) (interface{}, error) {
		return SomeFunc(params[0].(int64), params[1].(string))
	}, int64(12), "abc")

	// create a function map for easy key access
	m := make(functions.FuncMap)
	m["f1"] = f1

	// access your func by a key
	f, ok := m["f1"]
	if ok {
		// run your func
		_, _ = f.Run()
	}
}

// here could be anything useful, this is just an example:
func SomeFunc(a int64, b string) (interface{}, error) {
	println("called f1:", a, b)
	return nil, nil
}

```