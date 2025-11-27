package main

import (
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
  sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{0 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
