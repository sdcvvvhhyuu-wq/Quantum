package icmp_tunnel

import (
    "log"
    "net"
    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"
)

type ICMPTunnel struct {
    conn *icmp.PacketConn
}

func NewICMPTunnel() *ICMPTunnel {
    conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
    if err != nil {
        log.Printf("ICMP tunnel listen failed (root required): %v", err)
        return nil
    }
    return &ICMPTunnel{conn: conn}
}

func (t *ICMPTunnel) Start() {
    if t == nil { return }
    go func() {
        log.Println("ICMP covert tunnel started")
        buf := make([]byte, 1500)
        for {
            n, peer, err := t.conn.ReadFrom(buf)
            if err != nil { continue }
            msg, _ := icmp.ParseMessage(1, buf[:n])
            if msg != nil && msg.Type == ipv4.ICMPTypeEcho {
                reply := icmp.Message{Type: ipv4.ICMPTypeEchoReply, Code: 0, Body: msg.Body}
                b, _ := reply.Marshal(nil)
                t.conn.WriteTo(b, peer)
            }
        }
    }()
}
