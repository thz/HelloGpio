package gpio

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

// WaitPin waits for edge on a pin, calls the callback, and waits a bit to prevent chatter.
func WaitPin(pinName string, callback func()) {
	pin := gpioreg.ByName(pinName)
	pin.In(gpio.PullUp, gpio.BothEdges)
	for {
		pin.WaitForEdge(-1)
		val := pin.Read()
		if val == gpio.Low {
			callback()
		}
		time.Sleep(200 * time.Millisecond)
	}
}
