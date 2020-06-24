package components

import (
	"fmt"
	"testing"
)

type IUint32 uint32
func TestType(t *testing.T) {
	var id IUint32 = 2
	var id2 uint32 = 3
	id = IUint32(id2)
	var test = make(map[IUint32]string)
	if _,ok := test[id];ok{
		fmt.Println(id2)
	}
	fmt.Println(id , id2 , test[2])
}