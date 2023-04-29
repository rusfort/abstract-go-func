package functions

import (
	"context"
	"fmt"
)

type AbstractBaseFunc struct {
	ctx    context.Context
	self   interface{}
	params []interface{}
}

type IAbstractBaseFunc interface {
	Run(params ...interface{}) (interface{}, error)
	GetContext() context.Context
	GetSelf() interface{}
	GetParams() []interface{}
}

func NewAbstractBaseFunc(
	ctx context.Context,
	self interface{},
	params ...interface{},
) *AbstractBaseFunc {
	return &AbstractBaseFunc{
		ctx:    ctx,
		self:   self,
		params: params,
	}
}

func (a *AbstractBaseFunc) Run(params ...interface{}) (interface{}, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (a *AbstractBaseFunc) GetContext() context.Context {
	return a.ctx
}

func (a *AbstractBaseFunc) GetSelf() interface{} {
	return a.self
}

func (a *AbstractBaseFunc) GetParams() []interface{} {
	return a.params
}

type AbstractFunc struct {
	IAbstractBaseFunc
	operation func(params ...interface{}) (interface{}, error)
}

func NewAbstractFunc(operation func(params ...interface{}) (interface{}, error), params ...interface{}) *AbstractFunc {
	return &AbstractFunc{
		IAbstractBaseFunc: NewAbstractBaseFunc(nil, nil, params...),
		operation:         operation,
	}
}

func (a *AbstractFunc) Run(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		params = a.GetParams() //using default params
	}
	return a.operation(params...)
}

//type AbstractFunc func(params ...interface{}) (interface{}, error)
type AbstractFuncWithContext func(ctx context.Context, params ...interface{}) (interface{}, error)
type AbstractMethod func(self interface{}, params ...interface{}) (interface{}, error)
type AbstractMethodWithContext func(self interface{}, ctx context.Context, params ...interface{}) (interface{}, error)

type FuncMap map[string]IAbstractBaseFunc
