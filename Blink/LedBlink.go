package main

import (
	"time"

	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	edisonAdaptor := edison.NewEdisonAdaptor("edison")
	//sensor := gpio.NewAnalogSensorDriver(edisonAdaptor, "sensor", "0")
	led := gpio.NewLedDriver(edisonAdaptor, "led", "13")

	work := func() {
		for {
			led.On()
			led.Brightness(1)
			time.Sleep(time.Second * 2)
			led.Off()
			time.Sleep(time.Second * 2)
			fmt.Println("blink")
		}
		// 	gbot.On(sensor.Event("data"), func(data interface{}) {
		// 		brightness := uint8(
		// 			gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 4096), 0, 255),
		// 		)
		// 		fmt.Println("sensor", data)
		// 		fmt.Println("brightness", brightness)
		// 		led.Brightness(brightness)
		// 	})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{edisonAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
