package server

import (
	"asjade/logger"
	"fmt"
	"sync"
)

// NewFunc .
type NewFunc func(options ...Option) (Server, error)

var manager map[string]NewFunc
var m sync.Mutex

var servers map[string]Server
var sm sync.Mutex

func init() {
	manager = make(map[string]NewFunc)
	servers = make(map[string]Server)
}

// RegisteServer .
func RegisteServer(serverName string, newFunc NewFunc) {
	m.Lock()
	defer m.Unlock()
	if _, ok := manager[serverName]; !ok {
		manager[serverName] = newFunc
		logger.Infof("server '%s' registed", serverName)
	}
}

// GetServerByName .
func GetServerByName(serverName string) (Server, error) {
	if s, ok := servers[serverName]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("server '%s' not found", serverName)
}

// Init .
func Init(options ...Option) error {
	m.Lock()
	defer m.Unlock()
	sm.Lock()
	defer sm.Unlock()
	for serverName, fn := range manager {
		server, err := fn(options...)
		if err != nil {
			return err
		}
		servers[serverName] = server
	}
	return nil
}

// StartServers .
func StartServers(options ...RunOption) error {
	m.Lock()
	defer m.Unlock()
	for _, server := range servers {
		if err := server.Start(options...); err != nil {
			return err
		}
	}
	return nil
}
