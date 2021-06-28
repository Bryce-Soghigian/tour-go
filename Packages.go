package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // NG
    seed_result := rand.Seed(time.Now().Unix())

    // OK
    // Seed does not need return value
    // rand.Seed(time.Now().UnixNano())

    fmt.Println(rand.Int())
}