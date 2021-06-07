package main

import (
	"time"
)

func garbageWorker() {
	for {
		ti := 1 * time.Hour
		ts := 5 * time.Minute
		tl := time.Now().Add(-ti).Unix()
		for i := 0; i < len(qs); i++ {
			if qs[i].CreatedAt < tl {
				qs = append(qs[:i], qs[i+1:]...)
				i--
			}
		}
		time.Sleep(ts)
	}
}
