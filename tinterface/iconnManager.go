package tinterface

type ConnManager interface {
	Add(conn Connection)
	Del(conn Connection)
	Get(connId uint32) (Connection, error)
	Len() int
	ClearConn()
}
