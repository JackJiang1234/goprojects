package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string){
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func TestEventBusPublish(t *testing.T){
	bus := NewEventBus()
	bus.Subscribe("topc1", sub1)
	bus.Subscribe("topc1", sub2)

	bus.Publish("topc1", "test1", "test2")
	bus.Publish("topc1", "testa", "testb")
	time.Sleep(1 * time.Second)
}