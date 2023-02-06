package conv_test

import (
	"fmt"
	"testing"

	"github.com/srcio/go/conv"
)

func TestConv(t *testing.T) {
	var i int = 10
	intPtr := conv.Ptr(i)
	intVal := conv.Val(intPtr)
	fmt.Printf("intPtr: %v\n", intPtr)
	fmt.Printf("intVal: %v\n", intVal)

	s := []int{1, 2, 3, 4, 5}
	intPtrSlice := conv.PtrSlice(s)
	intValSlice := conv.ValSlice(intPtrSlice)
	fmt.Printf("intPtrSlice: %v\n", intPtrSlice)
	fmt.Printf("intValSlice: %v\n", intValSlice)

	m := map[string]string{"a": "1", "b": "2", "c": "3"}
	stringPtrMap := conv.PtrMap(m)
	stringValMap := conv.ValMap(stringPtrMap)
	fmt.Printf("stringPtrMap: %v\n", stringPtrMap)
	fmt.Printf("stringValMap: %v\n", stringValMap)
}
