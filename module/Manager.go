package module

import (
	"errors"
	"fmt"
	"reflect"
)

type Manager struct {
	modules     map[string]IModule     // 挂载的模块
	modulesFunc map[string]interface{} // 挂载模块的方法
}

var ErrModuleNotExists = errors.New("module not exists")
var ErrModuleFuncNotExists = errors.New("module func not exists")

// 初始化
func (mgr *Manager) InitManager() {
	fmt.Printf("InitManager \n")
	mgr.modules = make(map[string]IModule)
	mgr.modulesFunc = make(map[string]interface{})
}

// 注册模块
func (mgr *Manager) RegisterModule(name string, module IModule) {
	mgr.modules[name] = module
}

// 初始化模块
func (mgr *Manager) InitModule(s *SandBox) {
	for name, creator := range getInitModules() {
		mgr.RegisterModule(name, creator(s))
	}
}

// 启动模块
func (mgr *Manager) StartModule() error {
	for name, m := range mgr.modules {
		if err := m.Start(); err != nil {
			fmt.Printf("module[%v] start with err: %v\n", name, err)
			return err
		}
	}

	return nil
}

// 获取模块
func (mgr *Manager) GetModule(name string) IModule {
	m := mgr.modules[name]

	return m
}

// 注册模块方法
func (mgr *Manager) RegisterModuleFunc(name string, bindFunc interface{}) {
	mgr.modulesFunc[name] = bindFunc
}

// 获取模块方法
func (mgr *Manager) GetModuleFunc(name string) (interface{}, bool) {
	f, found := mgr.modulesFunc[name]

	return f, found
}

// 调用模块方法
func (mgr *Manager) CallModuleFunc(moduleName string, methodName string, args []reflect.Value, results ...interface{}) error {
	m := mgr.GetModule(moduleName)
	if m == nil {
		return ErrModuleNotExists
	}

	method, found := mgr.GetModuleFunc(methodName)
	if !found {
		return ErrModuleFuncNotExists
	}

	return m.Call(method, args, results...)
}
