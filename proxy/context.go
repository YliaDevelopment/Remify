package proxy

import (
	"net"

	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Context struct {
	Listener *minecraft.Listener

	Token  oauth2.TokenSource
	logger *zap.SugaredLogger

	ServerAddress string
	ListenAddress string
}

type (
	PacketFunc      func(header packet.Header, payload []byte, src, dst net.Addr)
	PacketCallback  func(pk packet.Packet, proxy *Context, toServer bool) (packet.Packet, error)
	ConnectCallback func(proxy *Context)
)
