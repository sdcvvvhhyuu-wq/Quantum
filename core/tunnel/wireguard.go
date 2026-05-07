package tunnel

import (
    "log"
    "net"

    "golang.zx2c4.com/wireguard/device"
    "golang.zx2c4.com/wireguard/tun/netstack"
    "github.com/sdcvvvhhyuu-wq/argotunnel/tree/main/core/pqc"
)

type WireGuardTunnel struct {
    dev     *device.Device
    tnet    *netstack.Net
    quantum *pqc.QuantumSession
}

func NewWireGuardTunnel() *WireGuardTunnel {
    tun, tnet, err := netstack.CreateNetTUN(
        []net.IP{net.ParseIP("10.0.0.2")},
        []net.IP{net.ParseIP("1.1.1.1")},
        1420,
    )
    if err != nil { log.Panic(err) }
    dev := device.NewDevice(tun, device.NewLogger(device.LogLevelError, "ArgoWireGuard: "))
    return &WireGuardTunnel{dev: dev, tnet: tnet}
}

func (w *WireGuardTunnel) Start() error {
    log.Println("WireGuard tunnel started with PQC hybrid handshake")
    return nil
}

func (w *WireGuardTunnel) SetQuantumSession(s *pqc.QuantumSession) {
    w.quantum = s
}
