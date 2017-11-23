package data

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type Memory struct {
	// ...
}

func (m Memory) Open(parm string) (io.ReadWriteCloser, error) {
	// ...
	// just for a example
	return nil, nil
}

func newMemoryStorage() Memory {
	return Memory{}
}

type Disk struct {
	// ...
}

func (d Disk) Open(parm string) (io.ReadWriteCloser, error) {
	// ...
	// just for a example
	return nil, nil
}

func newDiskStorage() Disk {
	return Disk{}
}

type Temp struct {
	// ...
}

func (m Temp) Open(parm string) (io.ReadWriteCloser, error) {
	// ...
	// just for a example
	return nil, nil
}

func newTempStorage() Temp {
	return Temp{}
}

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()
	}
}
