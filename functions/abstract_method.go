package functions

import (
	"context"
)

// abstract method of a given class

type AbstractMethod struct {
	IAbstractBaseFunc
	operation func(self interface{}, params ...interface{}) (interface{}, error)
}

func NewAbstractMethod(self interface{}, operation func(self interface{}, params ...interface{}) (interface{}, error), params ...interface{}) *AbstractMethod {
	return &AbstractMethod{
		IAbstractBaseFunc: NewAbstractBaseFunc(self, params...),
		operation:         operation,
	}
}

func (a *AbstractMethod) Run(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		params = a.GetParams() //using default params
	}
	return a.operation(a.GetSelf(), params...)
}

func (a *AbstractMethod) RunWithContext(_ context.Context, params ...interface{}) (interface{}, error) {
	return a.Run(params...)
}

// abstract method of a given class with context

type AbstractMethodWithContext struct {
	IAbstractBaseFunc
	operation func(self interface{}, ctx context.Context, params ...interface{}) (interface{}, error)
}

func NewAbstractMethodWithContext(self interface{}, operation func(self interface{}, ctx context.Context, params ...interface{}) (interface{}, error), params ...interface{}) *AbstractMethodWithContext {
	return &AbstractMethodWithContext{
		IAbstractBaseFunc: NewAbstractBaseFunc(self, params...),
		operation:         operation,
	}
}

func (a *AbstractMethodWithContext) RunWithContext(ctx context.Context, params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		params = a.GetParams() //using default params
	}
	return a.operation(a.GetSelf(), ctx, params...)
}
