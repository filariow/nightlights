package main

import (
	"machine"
	"time"
)

var (
	threshold      = uint16(40000)
	waitTime       = time.Minute * 10
	fetchRetry     = 6
	sampleInterval = time.Second * 10
)

func main() {
	machine.InitADC()
	pr := machine.ADC{machine.ADC0}

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	r := machine.D2
	r.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		v := fetchLight(pr)
		switchLight(led, r, v)
		time.Sleep(waitTime)
	}
}

func fetchLight(pr machine.ADC) uint16 {
	s := uint32(0)
	for i := 1; i <= fetchRetry; i++ {
		v := pr.Get()
		s += uint32(v)
		if i != fetchRetry {
			time.Sleep(sampleInterval)
		}
	}

	return uint16(s / uint32(fetchRetry))
}

func switchLight(led, r machine.Pin, light uint16) {
	if light > threshold {
		led.High()
		r.High()
		return
	}

	led.Low()
	r.Low()
}
