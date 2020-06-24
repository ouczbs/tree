package netutil

import (
	"encoding/binary"
	"net"
	"sync"
	"github.com/ouczbs/tree/engine/consts"
	"github.com/ouczbs/tree/engine/gwlog"
	"sync/atomic"
)
var (
	packetEndian               = binary.LittleEndian
	predefinePayloadCapacities []uint32

	debugInfo struct {
		NewCount     int64
		AllocCount   int64
		ReleaseCount int64
	}

	packetBufferPools = map[uint32]*sync.Pool{}
	packetPool        = sync.Pool{
		New: func() interface{} {
			p := &Packet{}
			p.bytes = p.initialBytes[:]

			if consts.DEBUG_PACKET_ALLOC {
				atomic.AddInt64(&debugInfo.NewCount, 1)
				gwlog.Infof("DEBUG PACKETS: ALLOC=%d, RELEASE=%d, NEW=%d",
					atomic.LoadInt64(&debugInfo.AllocCount),
					atomic.LoadInt64(&debugInfo.ReleaseCount),
					atomic.LoadInt64(&debugInfo.NewCount))
			}
			return p
		},
	}
)
// ConnectTCP connects to host:port in TCP
func ConnectTCP(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	return conn, err
}
func init() {

	if _MIN_PAYLOAD_CAP >= consts.PACKET_PAYLOAD_LEN_COMPRESS_THRESHOLD {
		gwlog.Fatalf("_MIN_PAYLOAD_CAP should be smaller than PACKET_PAYLOAD_LEN_COMPRESS_THRESHOLD")
	}

	payloadCap := uint32(_MIN_PAYLOAD_CAP) << _CAP_GROW_SHIFT
	for payloadCap < _MAX_PAYLOAD_LENGTH {
		predefinePayloadCapacities = append(predefinePayloadCapacities, payloadCap)
		payloadCap <<= _CAP_GROW_SHIFT
	}
	predefinePayloadCapacities = append(predefinePayloadCapacities, _MAX_PAYLOAD_LENGTH)

	for _, payloadCap := range predefinePayloadCapacities {
		payloadCap := payloadCap
		packetBufferPools[payloadCap] = &sync.Pool{
			New: func() interface{} {
				return make([]byte, _PREPAYLOAD_SIZE+payloadCap)
			},
		}
	}
}

func getPayloadCapOfPayloadLen(payloadLen uint32) uint32 {
	for _, payloadCap := range predefinePayloadCapacities {
		if payloadCap >= payloadLen {
			return payloadCap
		}
	}
	return _MAX_PAYLOAD_LENGTH
}



