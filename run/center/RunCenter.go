package main

import (
	"github.com/ouczbs/tree/components/center"
)
func main(){
	service := center.NewCenterService()
	service.Run()
}