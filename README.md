# module

## 模块框架代码

> 代码模块化设计，利用反射实现模块与模块之间解耦和

代码文件

### IModule.go
    
    IModule接口，包含3个函数
    
    ```
    Call    // 进行模块内对外方法调用
    Start   // 开始
    Stop    // 停止
    ```
    
    模块构造器Creator，对外一个初始化注册模块Init，业务代码调用把模块注册到内存中
    
    ```
    type Creator func(s *SandBox) IModule
    var initModules = make(map[string]Creator)
    // 初始化注册模块
    func Init(name string, creator Creator) {
        initModules[name] = creator
    }
    ```
    
### Base.go

    定义了一个Base结构，对接口IModule的具体现实，主要是现实了Call调用模块内对外函数方法 
    
    对外了一个GetValues方法，用来把参数处理为Call所需要的映射Value类型
  
### Manager.go
    
    定于了Manager结构体
    
    ```
    modules     map[string]IModule     // 挂载的模块
    modulesFunc map[string]interface{} // 挂载模块的方法
    ```
    
    主要对外功能
    
    ```
    InitManager：初始化Manager结构体
    InitModule：将initModules内的模块注册到Manager.modules 中
    StartModule：启动模块
    RegisterModuleFunc：注册模块对外方法
    CallModuleFunc：调用对应模块对应的方法
    ```
    
## 栗子

代码在example目录下
