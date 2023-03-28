package singleton

type Singleton interface {
	AddOne() int
}

type eagerSingleton struct {
	count int
}

func (s *eagerSingleton) AddOne() int {
	s.count++
	return s.count
}

var singleton *eagerSingleton

func init() {
	singleton = &eagerSingleton{}
}

func GetInstance() Singleton {
	return singleton
}
