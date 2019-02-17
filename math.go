package eapi

import (
    "math/rand"
    "time"
)

func RandomNumber(min, max int) int {
    rand.Seed(time.Now().Unix())
    if (max - min) < 0 {
        c := max
        max = min
        min = c
    }

    return rand.Intn(max-min) + min
}