package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	_ "gobot.io/x/gobot/platforms/firmata"
	"time"
)

func Arduino() *gobot.Robot {
	fa := firmata.NewAdaptor("/dev/ttyACM0")
	led := gpio.NewLedDriver(fa, "13")

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			led.Toggle()
		})
	}
	robot := gobot.NewRobot("Blinky",
		[]gobot.Connection{fa},
		[]gobot.Device{led},
		work)

	return robot
}

func main() {
	master := gobot.NewMaster()
	master.AddRobot(Arduino())

	fmt.Println("Starting master")
	master.Start()
}
