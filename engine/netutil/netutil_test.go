package netutil

import (
	"fmt"
	"testing"
)
func switchtype(fl Flushable){
	fl.Flush()
	q := fl.(* Connection)
	q.Flush()
	fmt.Println(fl , q)
}
func TestConnectTCP(t *testing.T) {
	p := &Connection{

	}
	switchtype(p)
}
