package main

import "time"

type ClockInterface interface {
    GetNow() time.Time
}

type Clock struct {}

func (c *Clock) GetNow() time.Time {
    return time.Now()
}
