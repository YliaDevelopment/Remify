package injectors

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Injector interface {
	OnLogin(client *minecraft.Conn, server *minecraft.Conn)
	OnClientPacket(clientPacket packet.Packet) (packet.Packet, error)
	OnServerPacket(serverPacket packet.Packet) (packet.Packet, error)
	Name() string
	Version() string
}
