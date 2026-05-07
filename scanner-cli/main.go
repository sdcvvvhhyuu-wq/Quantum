package main

import (
    "flag"
    "fmt"
    "net"
    "time"
)

func main() {
    cidr := flag.String("cidr", "104.16.0.0/12", "CIDR to scan")
    flag.Parse()
    _, ipnet, _ := net.ParseCIDR(*cidr)
    for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
        if isClean(ip) { fmt.Println(ip) }
    }
}

func inc(ip net.IP) {
    for j := len(ip)-1; j>=0; j-- {
        ip[j]++
        if ip[j] > 0 { break }
    }
}

func isClean(ip net.IP) bool {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip.String(), "443"), 500*time.Millisecond)
    if err == nil { conn.Close(); return true }
    return false
}
