package netutil

import (
	"encoding/binary"
	"fmt"
	"github.com/ouczbs/tree/engine/consts"
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/ioutil"
	"github.com/ouczbs/tree/engine/oputil"
	"github.com/pkg/errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	_MAX_PACKET_SIZE    = 25 * 1024 * 1024 // _MAX_PACKET_SIZE is the max size limit of packets in packet connections
	_SIZE_FIELD_SIZE    = 4                // _SIZE_FIELD_SIZE is the packet size field (uint32) size
	_PREPAYLOAD_SIZE    = _SIZE_FIELD_SIZE
	_PREPAYLOAD_RESRTVE = 2
	_MAX_PAYLOAD_LENGTH = _MAX_PACKET_SIZE - _PREPAYLOAD_SIZE
)

var (
	// NETWORK_ENDIAN is the network Endian of connections
	NETWORK_ENDIAN = binary.LittleEndian
	errRecvAgain   = _ErrRecvAgain{}
)

type _ErrRecvAgain struct{}

func (err _ErrRecvAgain) Error() string {
	return "recv again"
}

func (err _ErrRecvAgain) Temporary() bool {
	return true
}

func (err _ErrRecvAgain) Timeout() bool {
	return true
}

// PacketConnection is a connection that send and receive data packets upon a network stream connection
type PacketConnection struct {
	conn               Connection
	pendingPackets     []*Packet
	pendingPacketsLock sync.Mutex

	// buffers and infos for receiving a packet
	payloadLenBuf         [_SIZE_FIELD_SIZE]byte
	payloadLenBytesRecved int
	recvTotalPayloadLen   uint32
	recvedPayloadLen      uint32
	recvingPacket         *Packet
}

// NewPacketConnection creates a packet connection based on network connection
func NewPacketConnection(conn Connection) *PacketConnection {
	pc := &PacketConnection{
		conn: conn,
	}

	return pc
}

// NewPacket allocates a new packet (usually for sending)
func (pc *PacketConnection) NewPacket() *Packet {
	return allocPacket()
}

// SendPacket send packets to remote
func (pc *PacketConnection) SendPacket(packet *Packet) error {
	if consts.DEBUG_PACKETS {
		gwlog.Debugf("%s SEND PACKET %p: msgtype=%v, payload(%d)=%v", pc, packet,
			packetEndian.Uint16(packet.bytes[_PREPAYLOAD_SIZE:_PREPAYLOAD_SIZE+2]),
			packet.GetPayloadLen(),
			packet.bytes[_PREPAYLOAD_SIZE+2:_PREPAYLOAD_SIZE+packet.GetPayloadLen()])
	}
	if atomic.LoadInt64(&packet.refcount) <= 0 {
		gwlog.Panicf("sending packet with refcount=%d", packet.refcount)
	}

	packet.AddRefCount(1)
	pc.pendingPacketsLock.Lock()
	pc.pendingPackets = append(pc.pendingPackets, packet)
	pc.pendingPacketsLock.Unlock()
	return nil
}

// Flush connection writes
func (pc *PacketConnection) Flush(reason string) (err error) {
	pc.pendingPacketsLock.Lock()
	if len(pc.pendingPackets) == 0 { // no packets to send, common to happen, so handle efficiently
		pc.pendingPacketsLock.Unlock()
		return
	}
	packets := make([]*Packet, 0, len(pc.pendingPackets))
	packets, pc.pendingPackets = pc.pendingPackets, packets
	pc.pendingPacketsLock.Unlock()

	// flush should only be called in one goroutine
	op := oputil.StartOperation("FlushPackets-" + reason)
	defer op.Finish(time.Millisecond * 300)

	//var cw *flate.Writer

	if len(packets) == 1 {
		// only 1 packet to send, just send it directly, no need to use send buffer
		packet := packets[0]

		err = ioutil.WriteAll(pc.conn, packet.data())
		gwlog.Debugf(" send package  length = %d ; data = %s " ,packet.GetPayloadLen(),packet.bytes[_PREPAYLOAD_SIZE : _PREPAYLOAD_SIZE+packet.GetPayloadLen()])
		packet.Release()
		if err == nil {
			err = pc.conn.Flush()
		}
		return
	}

	for _, packet := range packets {
		ioutil.WriteAll(pc.conn, packet.data())
		gwlog.Debugf(" send package  length = %d ; data = %s " ,packet.GetPayloadLen(),packet.bytes[_PREPAYLOAD_SIZE : _PREPAYLOAD_SIZE+packet.GetPayloadLen()])
		packet.Release()
	}

	// now we send all data in the send buffer
	if err == nil {
		err = pc.conn.Flush()
	}
	return
}

// SetRecvDeadline sets the receive deadline
func (pc *PacketConnection) SetRecvDeadline(deadline time.Time) error {
	return pc.conn.SetReadDeadline(deadline)
}

// RecvPacket receives the next packet
func (pc *PacketConnection) RecvPacket() (*Packet, error) {
	if pc.payloadLenBytesRecved < _SIZE_FIELD_SIZE {
		// receive more of payload len bytes
		n, err := pc.conn.Read(pc.payloadLenBuf[pc.payloadLenBytesRecved:])
		pc.payloadLenBytesRecved += n
		if pc.payloadLenBytesRecved < _SIZE_FIELD_SIZE {
			if err == nil {
				err = errRecvAgain
			}
			return nil, err // packet not finished yet
		}

		pc.recvTotalPayloadLen = NETWORK_ENDIAN.Uint32(pc.payloadLenBuf[:])

		if pc.recvTotalPayloadLen == 0 || pc.recvTotalPayloadLen > _MAX_PAYLOAD_LENGTH {
			err := errors.Errorf("invalid payload length: %v", pc.recvTotalPayloadLen)
			pc.resetRecvStates()
			pc.Close()
			return nil, err
		}

		pc.recvedPayloadLen = 0
		pc.recvingPacket = NewPacket()
		pc.recvingPacket.AssureCapacity(pc.recvTotalPayloadLen)
	}

	// now all bytes of payload len is received, start receiving payload
	n, err := pc.conn.Read(pc.recvingPacket.bytes[_PREPAYLOAD_SIZE+pc.recvedPayloadLen : _PREPAYLOAD_SIZE+pc.recvTotalPayloadLen])
	pc.recvedPayloadLen += uint32(n)

	if pc.recvedPayloadLen == pc.recvTotalPayloadLen {
		// full packet received, return the packet
		packet := pc.recvingPacket
		packet.SetPayloadLen(pc.recvTotalPayloadLen)
		pc.resetRecvStates()

		return packet, nil
	}

	if err == nil {
		err = errRecvAgain
	}
	return nil, err
}
func (pc *PacketConnection) resetRecvStates() {
	pc.payloadLenBytesRecved = 0
	pc.recvTotalPayloadLen = 0
	pc.recvedPayloadLen = 0
	pc.recvingPacket = nil
}

// Close the connection
func (pc *PacketConnection) Close() error {
	return pc.conn.Close()
}

// RemoteAddr return the remote address
func (pc *PacketConnection) RemoteAddr() net.Addr {
	return pc.conn.RemoteAddr()
}

// LocalAddr returns the local address
func (pc *PacketConnection) LocalAddr() net.Addr {
	return pc.conn.LocalAddr()
}

func (pc *PacketConnection) String() string {
	return fmt.Sprintf("[%s >>> %s]", pc.LocalAddr(), pc.RemoteAddr())
}
