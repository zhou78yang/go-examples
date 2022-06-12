package main

import (
    "fmt"
    "os"
)

func main() {
    // 永远不会执行
    defer fmt.Println("exit!")

    os.Exit(3)
}
