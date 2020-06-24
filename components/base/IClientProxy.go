package base

import "github.com/ouczbs/tree/engine/proto"

type IClientProxy interface{
	proto.IRequestProxy
	Serve()
}
