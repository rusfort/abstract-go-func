# abstract-go-func

### Lib for abstract functions and methods in go
### Try this out!

## Example 1: Basic function with some parameters

```go
package main

import (
	"github.com/rusfort/abstract-go-func/functions"
)

func main() {
	println("hello")

	// create your new abstract function
	f1 := functions.NewAbstractFunc(func(params ...interface{}) (interface{}, error) {
		return SomeFunc(params[0].(int), params[1].(string)) // NOTE: if you feel that someone can make a mess with data types, assert it before return!
	}, 12, "abc") // set default params if needed

	// create a function map for easy key access
	m := make(functions.FuncMap)
	m["f1"] = f1

	// access your func by a key
	f, ok := m["f1"]
	if ok {
		// run your func with default params
		_, _ = f.Run()
		// run your func with custom params
		_, _ = f.Run(42, "qwerty")
	}
}

// here could be anything useful, this is just an example:
func SomeFunc(a int, b string) (interface{}, error) {
	println("called f1:", a, b)
	return nil, nil
}


```

## Example 2: Method of a given class with context and some parameters

```go
package main

import (
	"context"

	"github.com/rusfort/abstract-go-func/functions"
)

func main() {
	println("hello")

	someStruct := NewSomeStruct("my struct")
	ctx := context.Background()

	// create your new abstract method with context
	f2 := functions.NewAbstractMethodWithContext(someStruct, func(self interface{}, ctx context.Context, params ...interface{}) (interface{}, error) {
		return self.(*SomeStruct).SomeFunc(ctx, params[0].(int), params[1].(string)) // NOTE: if you feel that someone can make a mess with data types, assert it before return!
	}, 12, "abc") // set default params if needed

	// create a function map for easy key access
	m := make(functions.FuncMap)
	m["f2"] = f2

	// access your func by a key
	f, ok := m["f2"]
	if ok {
		// run your func with default params
		_, _ = f.RunWithContext(ctx)
		// run your func with custom params
		_, _ = f.RunWithContext(ctx, 42, "qwerty")
	}
}

type SomeStruct struct {
	name string
}

func NewSomeStruct(name string) *SomeStruct {
	return &SomeStruct{name: name}
}

// here could be anything useful, this is just an example:
func (s *SomeStruct) SomeFunc(ctx context.Context, a int, b string) (interface{}, error) {
	println("called f2:", s.name, a, b)
	return nil, nil
}

```