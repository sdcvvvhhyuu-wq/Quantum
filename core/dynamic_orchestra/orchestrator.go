package dynamic_orchestra

import (
    "log"
    "math/rand"
    "time"
    "github.com/argotunnel/core/tunnel"
)

type Orchestrator struct {
    current tunnel.Tunnel
    wg      tunnel.Tunnel
    ss      tunnel.Tunnel
}

func NewOrchestrator(wg, ss tunnel.Tunnel) *Orchestrator {
    return &Orchestrator{wg: wg, ss: ss, current: wg}
}

func (o *Orchestrator) Start() {
    go func() {
        for {
            time.Sleep(30 * time.Second)
            loss := rand.Float32()
            if loss > 0.2 {
                log.Println("Orchestrator: high loss, switching to Shadowsocks")
                o.SwitchTo(o.ss)
            } else {
                log.Println("Orchestrator: stable, keeping WireGuard")
                o.SwitchTo(o.wg)
            }
        }
    }()
}

func (o *Orchestrator) SwitchTo(newTun tunnel.Tunnel) {
    if o.current == newTun { return }
    o.current = newTun
    o.current.Start()
}
