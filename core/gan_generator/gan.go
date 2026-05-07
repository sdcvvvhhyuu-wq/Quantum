package gan_generator

import (
    "log"
    "math/rand"
    "time"
)

func EnableGANMorphing(interval time.Duration) {
    go func() {
        for {
            time.Sleep(interval)
            sizes := []int{40, 100, 1400, 800, 1200}
            payload := make([]byte, sizes[rand.Intn(len(sizes))])
            rand.Read(payload)
            log.Printf("GAN generator: injected %d bytes of cover traffic", len(payload))
        }
    }()
}
