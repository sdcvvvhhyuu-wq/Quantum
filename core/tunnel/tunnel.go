package tunnel

import "github.com/sdcvvvhhyuu-wq/argotunnel/tree/main/core/pqc"

type Tunnel interface {
    Start() error
    SetQuantumSession(s *pqc.QuantumSession)
}
