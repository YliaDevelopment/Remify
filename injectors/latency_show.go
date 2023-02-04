package injectors

import (
	"fmt"

	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type LatencyShow struct {
	client *minecraft.Conn
	server *minecraft.Conn
}

func (self *LatencyShow) OnLogin(client *minecraft.Conn, server *minecraft.Conn) {
	self.client = client
	self.server = server
}

func (self *LatencyShow) OnClientPacket(spacket packet.Packet) (packet.Packet, error) {
	self.client.WritePacket(&packet.Text{
		TextType: packet.TextTypeJukeboxPopup,
		Message:  fmt.Sprintf("%s | %s", self.server.Latency(), self.client.Latency()),
	})

	return spacket, nil
}

func (self *LatencyShow) OnServerPacket(spacket packet.Packet) (packet.Packet, error) {

	return spacket, nil
}
func (self *LatencyShow) Name() string {
	return "Latency show"
}

func (self *LatencyShow) Version() string {
	return "0.1.0"
}
