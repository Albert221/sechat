package domain

type Connection interface {
	Read() ([]byte, error)
	Write(data []byte) error
	CloseGracefully()
	Close()
}
