package tunnel

import (
    "log"
    "github.com/txthinking/socks5"
    "github.com/argotunnel/core/pqc"
)

type ShadowsocksTunnel struct {
    quantum *pqc.QuantumSession
}

func NewShadowsocksTunnel() *ShadowsocksTunnel { return &ShadowsocksTunnel{} }

func (s *ShadowsocksTunnel) Start() error {
    svr, _ := socks5.NewClassicServer(":0", "user", "secret", 5, 60)
    go svr.ListenAndServe(nil)
    log.Println("Shadowsocks tunnel started")
    return nil
}

func (s *ShadowsocksTunnel) SetQuantumSession(sess *pqc.QuantumSession) {
    s.quantum = sess
}
