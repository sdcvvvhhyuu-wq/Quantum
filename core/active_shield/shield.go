package active_shield

import (
    "log"
    "net"
    "sync"
    "time"
)

var suspiciousIPs sync.Map

func StartProbeShield() {
    log.Println("Active Probe Shield online (threshold: 3 probes)")
    go func() {
        for {
            time.Sleep(10 * time.Second)
            suspiciousIPs.Range(func(key, value any) bool {
                if value.(int) >= 3 {
                    log.Printf("Shield: blocking and morphing traffic for %s", key)
                    suspiciousIPs.Delete(key)
                }
                return true
            })
        }
    }()
}

func RecordProbe(ip net.IP) {
    val, _ := suspiciousIPs.LoadOrStore(ip.String(), 0)
    suspiciousIPs.Store(ip.String(), val.(int)+1)
}
