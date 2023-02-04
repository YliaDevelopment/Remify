package proxy

import (
	"fmt"
	"sync"

	"github.com/sandertv/gophertunnel/minecraft"
)

/*

	self.logger.Info("Starting game")

	go func() {
		if err := client.StartGame(server.GameData()); err != nil {
			self.logger.Errorf("%v", err)
		}
		self.logger.Info("-=")

	}()

	self.logger.Info("Spawning")
	if err := server.DoSpawn(); err != nil {
		self.logger.Errorf("%v", err)
	}
	self.logger.Info("Spawned!")


	for {
	}
*/

func prepareGame(client *minecraft.Conn, server *minecraft.Conn) error {
	var w sync.WaitGroup
	errs := make(chan error, 2)

	w.Add(1)
	go func() {
		defer w.Done()
		errs <- client.StartGame(server.GameData())
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		errs <- server.DoSpawn()
	}()

	w.Wait()

	for i := 0; i < 2; i++ {
		err := <-errs
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Context) handleClient(client *minecraft.Conn) error {
	defer client.Close()

	clientData := client.ClientData()

	self.logger.Infof("Accepted client: %v", clientData.ThirdPartyName)
	self.logger.Debug("Connecting to server")

	server, err := self.ConnectServer(false, clientData)

	if err != nil {
		return fmt.Errorf("failed to connect to server: %v", err)
	}

	self.Server = server

	defer server.Close()
	self.logger.Info("Connected!")

	self.logger.Info("Preparing game")

	if err := prepareGame(client, server); err != nil {
		return fmt.Errorf("failed to prepare game: %v", err)
	}

	self.logger.Info("Preparation done. Proxing packets!")
	var w sync.WaitGroup

	w.Add(1)

	go func() {
		defer w.Done()
		for {
			packet, err := server.ReadPacket()

			if err != nil {
				self.logger.Errorf("Error reading packet: %v", err)

				break
			}

			err = client.WritePacket(packet)

			if err != nil {
				self.logger.Errorf("Error writing packet: %v", err)

				break
			}
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		for {
			packet, err := client.ReadPacket()

			if err != nil {
				self.logger.Errorf("Error reading packet: %v", err)

				break
			}

			err = server.WritePacket(packet)

			if err != nil {
				self.logger.Errorf("Error writing packet: %v", err)

				break
			}
		}
	}()

	w.Wait()

	return nil
}
