package tnet

import (
	"errors"
	"sync"
	"tinynet/tinterface"
)

type ConnManager struct {
	connections map[uint32]tinterface.Connection //管理的连接信息
	connLock    sync.RWMutex                     //读写连接的读写锁
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]tinterface.Connection),
	}
}

func (c *ConnManager) Add(conn tinterface.Connection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
}

func (c *ConnManager) Del(conn tinterface.Connection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())
}

func (c *ConnManager) Get(connId uint32) (tinterface.Connection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
}
