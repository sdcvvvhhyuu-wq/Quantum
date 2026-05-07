package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tree/main/core/engine"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tree/main/core/capabilities"
)

func main() {
    fmt.Println("ArgoTunnel Ultimate – Linux")
    caps := capabilities.SelectOptimal()
    exec := engine.NewExecutor(caps)
    go exec.Start()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c
}
