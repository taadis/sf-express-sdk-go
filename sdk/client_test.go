package sdk

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	filename = filepath.Base(filename) // 获取当前文件名称: client_test.go
	fmt.Println(filename, " begin...")
	m.Run()
	fmt.Println(filename, " ...end")
}

// 测试新客户端
func TestNewClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Error(err)
	}
}

// 测试生成认证码函数
func TestComputeVerifyCode(t *testing.T) {
	verifyCode := computeVerifyCode("123456")
	if verifyCode != "4QrcOUm6Wau+VuBX8g+IPg==" {
		t.Fail()
	}
}
