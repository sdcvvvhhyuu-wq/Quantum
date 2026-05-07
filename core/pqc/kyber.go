package pqc

import (
    "crypto/rand"
    "log"
    "time"

    "github.com/cloudflare/circl/kem/kyber/kyber1024"
)

type QuantumSession struct {
    PrivateKey []byte
    PublicKey  []byte
    tunnel     interface{ SetQuantumSession(*QuantumSession) }
}

func NewQuantumSession(level int) (*QuantumSession, error) {
    pub, priv, err := kyber1024.GenerateKeyPair(rand.Reader)
    if err != nil {
        return nil, err
    }
    pubBytes, _ := pub.MarshalBinary()
    privBytes, _ := priv.MarshalBinary()
    return &QuantumSession{PrivateKey: privBytes, PublicKey: pubBytes}, nil
}

func (s *QuantumSession) StartKeyRollover() {
    go func() {
        for {
            time.Sleep(45 * time.Second)
            newPub, newPriv, err := kyber1024.GenerateKeyPair(rand.Reader)
            if err != nil {
                continue
            }
            pubBytes, _ := newPub.MarshalBinary()
            privBytes, _ := newPriv.MarshalBinary()
            s.PrivateKey = privBytes
            s.PublicKey = pubBytes
            log.Println("PQC key rotated (Kyber-1024 hybrid)")
            if s.tunnel != nil {
                s.tunnel.SetQuantumSession(s)
            }
        }
    }()
}

func (s *QuantumSession) RegisterTunnel(t interface{ SetQuantumSession(*QuantumSession) }) {
    s.tunnel = t
}

func (s *QuantumSession) HybridEncapsulate(classicPub []byte) (ct, ss []byte, err error) {
    scheme := kyber1024.Scheme()
    pk, err := scheme.UnmarshalBinaryPublicKey(s.PublicKey)
    if err != nil {
        return nil, nil, err
    }
    ct, ss, err = scheme.Encapsulate(pk)
    if err != nil {
        return nil, nil, err
    }
    hybrid := append(ss, classicPub...)
    return ct, hybrid, nil
}
