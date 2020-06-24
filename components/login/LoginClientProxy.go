package login

import (
	"net"

	"github.com/ouczbs/tree/engine/consts"
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/ioutil"
	"github.com/ouczbs/tree/engine/netutil"
	"github.com/ouczbs/tree/engine/post"
	"github.com/ouczbs/tree/engine/proto"
)

type loginClientProxy struct {
	*proto.PbConnection
	owner *loginService
	entityId TEntityId
}

func newLoginClientProxy(owner *loginService, conn net.Conn) *loginClientProxy {
	pbc := proto.NewPbConnection(*netutil.NewConnection(conn))
	proxy := &loginClientProxy{
		PbConnection: pbc,
		owner:        owner,
	}
	pbc.SetAutoFlush(consts.ENGINE_COMPONENT_WRITE_FLUSH_INTERVAL)
	return proxy
}

func (proxy *loginClientProxy) Serve() {
	// Serve the dispatcher client from server / gate
	defer func() {
		proxy.Close()
		post.Post(func() {
			proxy.owner.handClientDisconnect(proxy)
		})
		err := recover()
		if err != nil {
			gwlog.TraceError("Client %s paniced with error: %v", proxy, err)
		}
	}()
	gwlog.Infof("New dispatcher client: %s", proxy)
	for {
		var messageType proto.TMessageType
		pkt, err := proxy.Recv(&messageType)
		if err != nil {
			if ioutil.IsTimeoutError(err) {
				continue
			}
			gwlog.Panic(err)
		}
		proxy.owner.messageQueue <- proto.UMessage{Proxy: proxy, MessageType: messageType, Packet: pkt}
	}
}
