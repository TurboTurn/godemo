package gotest

import "testing"

//go test
func TestDivision(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("Division()测试通过")
	}
}
func TestDivision2(t *testing.T) {
	t.Error("就是不通过！")
}
