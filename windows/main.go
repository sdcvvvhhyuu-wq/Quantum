package main

import (
    "fmt"
    "os"
    "os/signal"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tree/main/core/engine"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tree/main/core/capabilities"
)

func main() {
    fmt.Println("ArgoTunnel Ultimate – Windows")
    caps := capabilities.SelectOptimal()
    exec := engine.NewExecutor(caps)
    go func() {
        if err := exec.Start(); err != nil {
            fmt.Println("Engine error:", err)
        }
    }()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}
