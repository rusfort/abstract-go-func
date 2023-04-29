package functions

import (
	"context"
	"fmt"
)

// abstract base class

type AbstractBaseFunc struct {
	ctx    context.Context
	self   interface{}
	params []interface{}
}

type IAbstractBaseFunc interface {
	Run(params ...interface{}) (interface{}, error)
	RunWithContext(ctx context.Context, params ...interface{}) (interface{}, error)
	GetSelf() interface{}
	GetParams() []interface{}
}

func NewAbstractBaseFunc(
	self interface{},
	params ...interface{},
) *AbstractBaseFunc {
	return &AbstractBaseFunc{
		self:   self,
		params: params,
	}
}

func (a *AbstractBaseFunc) Run(params ...interface{}) (interface{}, error) {
	panic(fmt.Errorf("cannot run without context"))
}

func (a *AbstractBaseFunc) RunWithContext(ctx context.Context, params ...interface{}) (interface{}, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (a *AbstractBaseFunc) GetSelf() interface{} {
	return a.self
}

func (a *AbstractBaseFunc) GetParams() []interface{} {
	return a.params
}

// map of abstract funcs
type FuncMap map[string]IAbstractBaseFunc
