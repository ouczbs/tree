package center

import (
	"github.com/ouczbs/tree/engine/consts"
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/ioutil"
	"github.com/ouczbs/tree/engine/netutil"
	"github.com/ouczbs/tree/engine/post"
	"github.com/ouczbs/tree/engine/proto"
	"net"
)

type UCenterClientProxy struct {
	*proto.PbConnection
	owner  *UCenterService
}

func newCenterClientProxy(owner *UCenterService, conn net.Conn) *UCenterClientProxy {
	pbc := proto.NewPbConnection(* netutil.NewConnection(conn))
	ccp := &UCenterClientProxy{
		PbConnection: pbc,
		owner:             owner,
	}
	pbc.SetAutoFlush(consts.ENGINE_COMPONENT_WRITE_FLUSH_INTERVAL)
	return ccp
}

func (ccp *UCenterClientProxy) Serve() {
	// Serve the dispatcher client from server / gate
	defer func() {
		ccp.Close()
		post.Post(func() {
			ccp.owner.handClientDisconnect(ccp)
		})
		err := recover()
		if err != nil {
			gwlog.TraceError("Client %s paniced with error: %v", ccp, err)
		}
	}()
	gwlog.Infof("New dispatcher client: %s", ccp)
	for {
		var messageType proto.TMessageType
		pkt, err := ccp.Recv(&messageType)
		if err != nil {
			if ioutil.IsTimeoutError(err) {
				continue
			}
			gwlog.Panic(err)
		}
		ccp.owner.messageQueue <- proto.UMessage{Proxy: ccp ,MessageType: messageType, Packet: pkt}
	}
}