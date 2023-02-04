package proxy

import (
	f "fmt"

	"github.com/mattn/go-colorable"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type DummyProto struct {
	id  int32
	ver string
}

func (p DummyProto) ID() int32            { return p.id }
func (p DummyProto) Ver() string          { return p.ver }
func (p DummyProto) Packets() packet.Pool { return packet.NewPool() }
func (p DummyProto) ConvertToLatest(pk packet.Packet, _ *minecraft.Conn) []packet.Packet {
	return []packet.Packet{pk}
}

func (p DummyProto) ConvertFromLatest(pk packet.Packet, _ *minecraft.Conn) []packet.Packet {
	return []packet.Packet{pk}
}

func (self *Context) Start() error {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	self.logger = logger.Sugar()

	logConf := zap.NewDevelopmentEncoderConfig()
	logConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	self.logger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(logConf),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)).Sugar()

	self.logger.Debug("Running initial connection...")
	initialServerConn, err := self.ConnectServer(true, login.ClientData{})

	if err != nil {
		return err
	}

	rr := initialServerConn.ResourcePacks()
	initialServerConn.Close()

	self.logger.Infof("Connected to server! Found %d resources.", len(rr))

	self.Listener, err = minecraft.ListenConfig{
		StatusProvider:       minecraft.NewStatusProvider("Remify proxy"),
		ResourcePacks:        rr,
		TexturePacksRequired: len(rr) > 1,
	}.Listen("raknet", self.ListenAddress)

	if err != nil {
		return f.Errorf("failed to start listening: %v", err)
	}

	self.logger.Infof("Listening on %s", self.ListenAddress)

	for {
		rawConn, err := self.Listener.Accept()

		if err != nil {
			return f.Errorf("failed to accept client: %v", err)
		}

		conn := rawConn.(*minecraft.Conn)

		go func() {
			err := self.handleClient(conn)
			if err != nil {
				self.logger.Errorf("failed to handle client: %v", err)
			}
		}()
	}
}
