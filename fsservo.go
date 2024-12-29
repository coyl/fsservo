package fsservo

import (
	"tinygo.org/x/drivers/servo"
)

const maxCW int16 = 1083
const minCW int16 = 1163
const stop int16 = 1168 // min 1164 max 1171
const maxCCW int16 = 1247
const minCCW int16 = 1172

func Stop(srv servo.Servo) {
	srv.SetMicroseconds(stop)
}

func Clockwise(srv servo.Servo, speed int16) {
	if speed < 0 || speed > 10 {
		speed = 10
	}
	speedTime := minCW + (maxCW-minCW)*speed/10
	srv.SetMicroseconds(speedTime)
}

func CounterClockwise(srv servo.Servo, speed int16) {
	if speed < 0 || speed > 10 {
		speed = 10
	}
	speedTime := minCCW + (maxCCW-minCCW)*speed/10
	srv.SetMicroseconds(speedTime)
}
