package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "github.com/argotunnel/core/engine"
    "github.com/argotunnel/core/capabilities"
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
