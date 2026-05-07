package capabilities

import (
    "math/rand"
    "strings"
)

type CapabilitySet struct {
    ActiveIDs      []int
    FragSize       int
    UseUTLS        bool
    UTLSProfile    string
    PQCLevel       int
    Transport      string
    AIMorphEnabled bool
    UseDNSTunnel   bool
    UseICMPTunnel  bool
    UseQRNG        bool
    DPIAdaptive    bool
}

func SelectOptimal() CapabilitySet {
    all := GetAllCapabilities()
    selected := CapabilitySet{ActiveIDs: []int{}}
    for _, c := range all {
        if rand.Float32() < 0.1 {
            selected.ActiveIDs = append(selected.ActiveIDs, c.ID)
            switch c.Category {
            case "frag":
                size := parseParamInt(c.Param, "size=")
                if size > 0 { selected.FragSize = size }
            case "tls":
                if strings.HasPrefix(c.Name, "utls_") {
                    selected.UseUTLS = true
                    selected.UTLSProfile = c.Param
                }
            case "pqc":
                level := parseParamInt(c.Param, "level=")
                if level > 0 { selected.PQCLevel = level }
            case "transport":
                selected.Transport = c.Param
            case "ai_evasion":
                selected.AIMorphEnabled = true
            case "dns":
                selected.UseDNSTunnel = true
            case "icmp":
                selected.UseICMPTunnel = true
            case "qrng":
                selected.UseQRNG = true
            case "dpi_analyzer":
                selected.DPIAdaptive = true
            }
        }
    }
    if selected.FragSize == 0 { selected.FragSize = 150 }
    if selected.PQCLevel == 0 { selected.PQCLevel = 2 }
    if selected.Transport == "" { selected.Transport = "wireguard" }
    return selected
}
