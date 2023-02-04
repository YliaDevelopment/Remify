package proxy

import (
	"github.com/YliaDevelopment/Remify/injectors"
	"github.com/sandertv/gophertunnel/minecraft"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Context struct {
	Listener *minecraft.Listener

	Token            oauth2.TokenSource
	logger           *zap.SugaredLogger
	EnabledInjectors []injectors.Injector

	ServerAddress string
	ListenAddress string
}
