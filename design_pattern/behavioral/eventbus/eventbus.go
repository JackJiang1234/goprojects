package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type Eventbus interface {
	Subscribe(topic string, hanle interface{}) error
	Publish(topic string, args ...interface{})
}

type asynEventBus struct {
	handler map[string][]reflect.Value
	sync.Mutex
}

func (bus *asynEventBus) Subscribe(topic string, f interface{}) error {
	bus.Lock()
	defer bus.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.handler[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handler[topic] = handler
	
	return nil
}

func (bus *asynEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handler[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}
	for i := range handlers {
		go handlers[i].Call(params)
	}
}

func NewEventBus() Eventbus {
	return &asynEventBus{
		handler: map[string][]reflect.Value{},
	}
}
