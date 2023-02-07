package xbdb

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) { //不能使用Testxbdb 类似名称，邪门
	rb := JoinBytes([]byte("ddd"), []byte("fff"))
	fmt.Println(string(rb))
}
