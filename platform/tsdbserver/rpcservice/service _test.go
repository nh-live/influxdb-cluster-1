package rpcservice_test

import (
	"fmt"
	"testing"
)

type f func()

func TestRpc_Open(t *testing.T) {
	var fn f
	fn = func() {

	}
	b(fn)
}

func b(a interface{}) {
	switch a.(type) {
	case f:
		fmt.Println("ok")
	}
	if a == nil {
		fmt.Println("ok")
	}
}
