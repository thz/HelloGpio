package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	hello_api "HelloGpio/lib/api"
	hello_gpio "HelloGpio/lib/gpio"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

// #########################################################33

var (
	pin_rot   gpio.PinIO
	pin_gelb  gpio.PinIO
	pin_gruen gpio.PinIO
)

func prepare() {
	// find pins
	pin_rot = gpioreg.ByName("GPIO21")
	pin_gelb = gpioreg.ByName("GPIO20")
	pin_gruen = gpioreg.ByName("GPIO16")

	// handle termination
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		rot(0)
		gelb(0)
		gruen(0)
		os.Exit(0)
	}()
}

// ###########################################################
func rot(an_aus int) {
	pin_rot.Out(an_aus != 0)
}

func gelb(an_aus int) {
	pin_gelb.Out(an_aus != 0)
}

func gruen(an_aus int) {
	pin_gruen.Out(an_aus != 0)
}

func alle(rot_an_aus, gelb_an_aus, gruen_an_aus int) {
	rot(rot_an_aus)
	gelb(gelb_an_aus)
	gruen(gruen_an_aus)
}

// Hello is an example callback function
func Hello() {
	fmt.Printf("Hello!\n")
}

func warte(period float64) {
	time.Sleep(time.Duration(1000*period) * time.Millisecond)
}

// Ampel loops through all states
func Ampel() {
	for {
		alle(1, 0, 0)
		warte(2)

		alle(1, 1, 0)
		warte(.5)

		alle(0, 0, 1)
		warte(2)

		alle(0, 1, 0)
		warte(.5)
	}
}

func roll() {
	if err := hello_api.Apicall("/example/uri", hello_api.API_VAL_UP); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	if err := hello_api.Apicall("/example/uri", hello_api.API_VAL_STOP); err != nil {
		panic(err)
	}
}

func main() {
	host.Init()
	prepare()

	roll()

	go Ampel()

	hello_gpio.WaitPin("GPIO5", Hello)
}
