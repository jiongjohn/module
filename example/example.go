package main

import (
	"fmt"
	"module/example/person"
	"module/module"
)

func main() {
	// 初始化模块加载
	sandbox := &module.SandBox{}
	sandbox.Manager.InitManager()

	sandbox.Manager.InitModule(sandbox)
	// 启动各模块的start
	sandbox.Manager.StartModule()

	sandbox.Manager.CallModuleFunc(person.MODULE_PERSON, person.MF_Person_SetName, module.GetValues("hello world"), nil)
	var personName string
	sandbox.Manager.CallModuleFunc(person.MODULE_PERSON, person.MF_Person_GetName, nil, &personName)
	fmt.Printf("personName:%v\n", personName)
}
