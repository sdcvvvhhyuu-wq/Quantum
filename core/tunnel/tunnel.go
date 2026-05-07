package tunnel

import "github.com/argotunnel/core/pqc"

type Tunnel interface {
    Start() error
    SetQuantumSession(s *pqc.QuantumSession)
}
