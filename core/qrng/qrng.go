package qrng

import (
    "crypto/rand"
    "encoding/binary"
    "math/big"
)

func RandInt(max int) int {
    n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
    return int(n.Int64())
}

func RandFloat() float64 {
    var b [8]byte
    rand.Read(b[:])
    return float64(binary.LittleEndian.Uint64(b[:])) / (1 << 64)
}
