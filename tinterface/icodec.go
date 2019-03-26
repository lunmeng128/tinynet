package tinterface

type Codec interface {
	UnPack(data []byte) (Message, error)
	Pack(msg Message) ([]byte, error)
}
