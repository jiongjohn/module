package module

import (
	"fmt"
	"sort"
	"testing"
)

var num []int

func init() {
	for n := 1; n < 200000; n++ {
		num = append(num, n)
	}
}

type TestModule struct {
	Base
}

func NewTestModule(s *SandBox) IModule {
	return &TestModule{}
}

func NewTestModule2(s *SandBox) *TestModule {
	return &TestModule{}
}

func (m *TestModule) GetBenchTest(a ...int) int {
	var sum int
	for i := range a {
		sum += a[i]
	}
	return sum
}

func (m *TestModule) TestReflectParam() (string, string) {
	return "heloasdojj", "大好dfa机会"
}

func (m *TestModule) TestReflectParamWithArgs(args ...int) []int {
	return args
}

func (m *TestModule) TestBool(a, b int) bool {
	return a == b
}

func (m *TestModule) TestInt(a, b int) int {
	s := a + b
	return s
}
func (m *TestModule) TestInt8(a, b int8) int8 {
	s := a + b
	return s
}

func (m *TestModule) TestArray(a, b int8) [2]int8 {
	s := [2]int8{a, b}
	return s
}

func (m *TestModule) TestSlice(a ...string) []string {
	var s []string
	s = append(s, a...)
	return s
}

func (m *TestModule) TestSlice2(a ...int) []int {
	var s []int
	s = append(s, a...)
	return s
}

type A struct {
	nums []int
}

func (m *TestModule) GetBenchTestArr() []int {
	nums := []int{3, 1, 2}
	sort.Ints(nums)
	return nums
}

func TestManager_CallModuleFunc(t *testing.T) {
	tm := &TestModule{}
	mgr := new(Manager)
	mgr.InitManager()
	Init("test", NewTestModule)
	mgr.InitModule(nil)
	mgr.StartModule()
	// 注册方法
	mgr.RegisterModuleFunc("test_func", tm.GetBenchTest)

	var resInt int
	err := mgr.CallModuleFunc("test", "test_func", GetValues(1, 2, 3), &resInt)
	if err != nil {
		t.Errorf("resInt add error  err, err：%v", err)
	}
	if resInt != 6 {
		t.Errorf("resInt is error: %v  right_res:%d", resInt, 6)
	}
	fmt.Printf("resInt =================== %v\n", resInt)

	// 测试数组
	a := &A{}
	mgr.RegisterModuleFunc("test_arr", tm.GetBenchTestArr)
	err = mgr.CallModuleFunc("test", "test_arr", nil, &a.nums)
	if err != nil {
		t.Errorf("test_arr add error err, err：%v", err)
	}
	fmt.Printf("test_arr =================== %v\n", a.nums)

	// 返回两个参数
	mgr.RegisterModuleFunc("test_str", tm.TestReflectParam)
	var resStr1, resStr2 string
	err = mgr.CallModuleFunc("test", "test_str", nil, &resStr1, &resStr2)
	if err != nil {
		t.Errorf("resStr add error  err, err：%v", err)
	}
	fmt.Printf("resStr =================== %v， %v\n", resStr1, resStr2)
}

// 利用基准测试测试历用反射获取方法
func BenchmarkUseReflect(b *testing.B) {
	tm := &TestModule{}
	mgr := new(Manager)
	mgr.InitManager()
	Init("test", NewTestModule)
	mgr.InitModule(nil)
	mgr.StartModule()
	// 注册方法
	mgr.RegisterModuleFunc("test_func", tm.GetBenchTest)

	for i := 0; i < b.N; i++ {
		var resInt int
		mgr.CallModuleFunc("test", "test_func", GetValues(1, 2, 3), &resInt)
	}
}

// 利用基准测试测试历不使用反射获取方法
func BenchmarkNoUseReflect(b *testing.B) {
	mgr := new(Manager)
	mgr.InitManager()
	mgr.RegisterModule("test", NewTestModule2(nil))
	mgr.StartModule()

	for i := 0; i < b.N; i++ {
		tm := mgr.GetModule("test").(*TestModule)
		tm.GetBenchTest(num...)
	}
}

// 常规方法
func BenchmarkNormal(b *testing.B) {
	m := NewTestModule2(nil)

	for i := 0; i < b.N; i++ {
		m.GetBenchTest(num...)
	}
}
