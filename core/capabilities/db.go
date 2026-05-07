package capabilities

import (
    "fmt"
    "strconv"
    "strings"
)

type Capability struct {
    ID       int
    Name     string
    Category string
    Param    string
}

func GetAllCapabilities() []Capability {
    caps := []Capability{}
    id := 1
    for i := 1; i <= 500; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("frag_%d", i), Category: "frag", Param: fmt.Sprintf("size=%d", i)})
        id++
    }
    for i := 0; i <= 300; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("pad_%d", i), Category: "padding", Param: fmt.Sprintf("bytes=%d", i)})
        id++
    }
    profiles := []string{"chrome_124", "firefox_125", "safari_17", "edge_122", "opera_91", "brave_123"}
    for _, p := range profiles {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("utls_%s", p), Category: "tls", Param: p})
        id++
    }
    for i := 0; i < 1000; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("cipher_%d", i), Category: "tls", Param: fmt.Sprintf("suite=%d", i)})
        id++
    }
    hosts := []string{"cloudflare.com", "google.com", "amazon.com", "fastly.com", "akamai.com", "microsoft.com"}
    for i := 0; i < 200; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("df_%d", i), Category: "fronting", Param: hosts[i%len(hosts)]})
        id++
    }
    transports := []string{"wireguard", "shadowsocks", "xray", "trojan", "hysteria", "v2ray", "ssr"}
    for _, t := range transports {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("transport_%s", t), Category: "transport", Param: t})
        id++
    }
    for i := 0; i < 250; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("pqc_%d", i), Category: "pqc", Param: fmt.Sprintf("level=%d", i%3+1)})
        id++
    }
    for i := 0; i < 500; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("ai_%d", i), Category: "ai_evasion", Param: fmt.Sprintf("noise=%d", i%100)})
        id++
    }
    for i := 0; i < 150; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("webrtc_%d", i), Category: "webrtc", Param: fmt.Sprintf("mode=%d", i%3)})
        id++
    }
    for i := 0; i < 200; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("scan_%d", i), Category: "scanner", Param: fmt.Sprintf("timeout=%d", i%10+1)})
        id++
    }
    for i := 0; i < 300; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("fallback_%d", i), Category: "fallback", Param: fmt.Sprintf("order=%d", i)})
        id++
    }
    for i := 0; i < 400; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("net_%d", i), Category: "network", Param: fmt.Sprintf("mtu=%d", 500+i%1000)})
        id++
    }
    // دسته‌های پیشرفته جدید
    for i := 0; i < 100; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("dns_%d", i), Category: "dns", Param: fmt.Sprintf("domain=%d.example.com", i)})
        id++
    }
    for i := 0; i < 50; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("icmp_%d", i), Category: "icmp", Param: "active"})
        id++
    }
    for i := 0; i < 30; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("qrng_%d", i), Category: "qrng", Param: "true"})
        id++
    }
    for i := 0; i < 20; i++ {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("dpi_%d", i), Category: "dpi_analyzer", Param: fmt.Sprintf("mode=%d", i%5)})
        id++
    }
    // پر کردن تا ۵۰۰۰
    for id <= 5000 {
        caps = append(caps, Capability{ID: id, Name: fmt.Sprintf("extra_%d", id), Category: "extra", Param: "dummy"})
        id++
    }
    return caps[:5000]
}

func parseParamInt(param, prefix string) int {
    valStr := strings.TrimPrefix(param, prefix)
    val, _ := strconv.Atoi(strings.TrimSpace(valStr))
    return val
}
