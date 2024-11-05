package main

import (
    "math/rand"
    "time"
)

const (
    N_ITERS = 1000
    TOLERANCE = 1e-5
)

func getObservation() float64 {
    time.Sleep(1*time.Nanosecond)
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()
}