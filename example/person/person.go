package person

import (
	"module/module"
)

// 模块别名
const MODULE_PERSON string = "PersonModule"

// 注册对外方法
const (
	MF_Person_GetName string = "MF_Person_GetName"
	MF_Person_SetName string = "MF_Person_SetName"
)

type PersonModule struct {
	module.Base
	sandbox *module.SandBox
	name    string
}

func NewPersonModule(s *module.SandBox) module.IModule {
	return &PersonModule{
		sandbox: s,
	}
}

func (m *PersonModule) Start() error {
	m.sandbox.RegisterModuleFunc(MF_Person_GetName, m.GetName)
	m.sandbox.RegisterModuleFunc(MF_Person_SetName, m.SetName)
	return nil
}

func (m *PersonModule) GetName() string {
	return m.name
}

func (m *PersonModule) SetName(name string) {
	m.name = name
}

func init() {
	module.Init(MODULE_PERSON, NewPersonModule)
}
