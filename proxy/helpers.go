package proxy

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (self *Context) ConnectServer(wantsRR bool, loginData login.ClientData) (*minecraft.Conn, error) {
	dialer := minecraft.Dialer{
		ClientData:  loginData,
		TokenSource: self.Token,
		DownloadResourcePack: func(id uuid.UUID, version string) bool {
			return wantsRR
		},
	}

	conn, err := dialer.Dial("raknet", self.ServerAddress)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
