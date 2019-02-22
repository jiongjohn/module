package module

import (
	"reflect"
)

type IModule interface {
	Call(method interface{}, args []reflect.Value, results ...interface{}) error
	Start() error // 开始
	Stop() error  // 停止
}

type Creator func(sandbox *SandBox) IModule

var initModules = make(map[string]Creator)

// 初始化注册模块
func Init(name string, creator Creator) {
	initModules[name] = creator
}

// 获取所有初始化注册的模块
func getInitModules() map[string]Creator {
	return initModules
}
