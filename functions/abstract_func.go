package functions

import (
	"context"
)

// abstract func

type AbstractFunc struct {
	IAbstractBaseFunc
	operation func(params ...interface{}) (interface{}, error)
}

func NewAbstractFunc(operation func(params ...interface{}) (interface{}, error), params ...interface{}) *AbstractFunc {
	return &AbstractFunc{
		IAbstractBaseFunc: NewAbstractBaseFunc(nil, params...),
		operation:         operation,
	}
}

func (a *AbstractFunc) Run(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		params = a.GetParams() //using default params
	}
	return a.operation(params...)
}

func (a *AbstractFunc) RunWithContext(_ context.Context, params ...interface{}) (interface{}, error) {
	return a.Run(params...)
}

// abstract func with context

type AbstractFuncWithContext struct {
	IAbstractBaseFunc
	operation func(ctx context.Context, params ...interface{}) (interface{}, error)
}

func NewAbstractFuncWithContext(operation func(ctx context.Context, params ...interface{}) (interface{}, error), params ...interface{}) *AbstractFuncWithContext {
	return &AbstractFuncWithContext{
		IAbstractBaseFunc: NewAbstractBaseFunc(nil, params...),
		operation:         operation,
	}
}

func (a *AbstractFuncWithContext) RunWithContext(ctx context.Context, params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		params = a.GetParams() //using default params
	}
	return a.operation(ctx, params...)
}
