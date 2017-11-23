package singleton

import "sync"

type singleton struct {
	N int
}

var (
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = singleton{}
	})

	return instance
}
