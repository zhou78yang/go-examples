package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("begin", time.Now())
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res, time.Now())
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1", time.Now())
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res, time.Now())
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2", time.Now())
    }
    fmt.Println("end", time.Now())
}
