package docs

import (
	"github.com/coyl/fsservo"
	"machine"
	"time"
	"tinygo.org/x/drivers/servo"
)

func main() {
	// Configure the servo motor
	servopwm := machine.Timer1
	servopwm.Configure(machine.PWMConfig{
		Period: 20e6,
	})
	servo1, _ := servo.New(servopwm, machine.D9)

	// Stop the servo motor
	fsservo.Stop(servo1)
	time.Sleep(time.Millisecond * 1000)

	for {
		// Rotate the servo motor clockwise with increasing speed
		for i := 1; i < 11; i++ {
			fsservo.Clockwise(servo1, int16(i))
			time.Sleep(time.Millisecond * 1000)
		}

		// Stop the servo motor
		fsservo.Stop(servo1)
		time.Sleep(time.Millisecond * 1000)

		// Rotate the servo motor counter-clockwise with increasing speed
		for i := 1; i < 11; i++ {
			fsservo.CounterClockwise(servo1, int16(i))
			time.Sleep(time.Millisecond * 1000)
		}

		// Stop the servo motor
		fsservo.Stop(servo1)
		time.Sleep(time.Millisecond * 1000)
	}
}
