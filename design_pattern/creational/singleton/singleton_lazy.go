package singleton

import (
	"sync"
)

var (
	lazySingletonInstance *lazySingleton
	once                  = &sync.Once{}
)

type lazySingleton struct {
}

func GetLazyInstance() *lazySingleton {
	if lazySingletonInstance == nil {
		once.Do(func() {
			lazySingletonInstance = &lazySingleton{}
		})
	}
	return lazySingletonInstance
}

