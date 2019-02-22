// 基础模块
package module

import (
	"errors"
	"reflect"
)

var (
	ErrCallNotFunc    = errors.New("funcInter is not func")
	ErrBindDestNotPtr = errors.New("bind dest is not ptr")
	ErrBindNoSettable = errors.New("bind non-settable variable passed to bind")
)

type Base struct {
}

func (m *Base) Start() error {
	return nil
}

func (m *Base) Stop() error {
	return nil
}

func (m *Base) RegFunc(name string, bindFunc interface{}) error {

	return nil
}

func (m *Base) Call(funcInter interface{}, paramsValue []reflect.Value, bind ...interface{}) error {
	v := reflect.ValueOf(funcInter)
	if v.Kind() != reflect.Func {
		return ErrCallNotFunc
	}

	values := v.Call(paramsValue) //方法调用并返回值
	for i := range values {
		if len(bind)-1 >= i {
			err := m.bind(bind[i], values[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Base) bind(dest interface{}, data reflect.Value) error {
	value := reflect.ValueOf(dest)
	if value.Kind() != reflect.Ptr {
		return ErrBindDestNotPtr
	}
	value = value.Elem()
	if !value.CanSet() {
		return ErrBindNoSettable
	}
	value.Set(data)
	return nil
}

func GetValues(param ...interface{}) []reflect.Value {
	values := make([]reflect.Value, 0, len(param))
	for i := range param {
		values = append(values, reflect.ValueOf(param[i]))
	}
	return values
}
