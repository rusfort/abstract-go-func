package functions

import (
	"context"
	_ "encoding/json"
	_ "fmt"
)

type AbstractFunc func(params ...interface{}) (interface{}, error)
type AbstractFuncWithContext func(ctx context.Context, params ...interface{}) (interface{}, error)
type AbstractMethod func(self interface{}, params ...interface{}) (interface{}, error)
type AbstractMethodWithContext func(self interface{}, ctx context.Context, params ...interface{}) (interface{}, error)

type FuncMap map[string]any
