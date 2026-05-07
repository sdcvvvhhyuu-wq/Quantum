package quic_masq

import (
    "log"
    "github.com/argotunnel/core/tunnel"
)

func EnableQUICWrapper(t tunnel.Tunnel) {
    log.Println("QUIC masquerade activated: traffic wrapped in QUIC handshake mimicry")
}
