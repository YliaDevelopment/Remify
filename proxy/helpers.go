package proxy

import (
	"time"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (self *Context) ConnectServer(wantsRR bool, loginData login.ClientData, timeout time.Duration) (*minecraft.Conn, error) {
	dialer := minecraft.Dialer{
		ClientData:  loginData,
		TokenSource: self.Token,
		DownloadResourcePack: func(id uuid.UUID, version string, current, total int) bool {
			return wantsRR
		},
	}

	conn, err := dialer.DialTimeout("raknet", self.ServerAddress, timeout)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
