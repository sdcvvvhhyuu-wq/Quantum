package obfs

import "log"

func EnableFragmentation(size int) {
    log.Printf("Fragmentation enabled: size %d bytes", size)
}
