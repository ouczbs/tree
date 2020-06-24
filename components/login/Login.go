package login

import "github.com/ouczbs/tree/components/base"

type(
	Component string
	ComponentId uint32
	TEntityId = uint32
	IClientProxy = base.IClientProxy
)

var (
	loginList = make(map[ComponentId]Component)
	gateList = make(map[ComponentId]Component)
	gameList = make(map[ComponentId]Component)

	dispatcherProxyList []IClientProxy
	centerProxy IClientProxy

)