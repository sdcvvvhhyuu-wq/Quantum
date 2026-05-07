package engine

import (
    "log"
    "time"

    "github.com/sdcvvvhhyuu-wq/argotunnel/core/capabilities"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/tunnel"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/obfs"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/pqc"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/ai_morph"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/active_shield"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/dynamic_orchestra"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/dns_tunnel"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/icmp_tunnel"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/dpi_analyzer"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/qrng"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/rl_agent"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/gan_generator"
    "github.com/sdcvvvhhyuu-wq/argotunnel/core/quic_masq"
)

type Executor struct {
    caps capabilities.CapabilitySet
}

func NewExecutor(c capabilities.CapabilitySet) *Executor {
    return &Executor{caps: c}
}

func (e *Executor) Start() error {
    log.Printf("Starting ArgoTunnel with %d active capabilities", len(e.caps.ActiveIDs))

    var primary, secondary tunnel.Tunnel
    switch e.caps.Transport {
    case "shadowsocks":
        primary = tunnel.NewShadowsocksTunnel()
        secondary = tunnel.NewWireGuardTunnel()
    default:
        primary = tunnel.NewWireGuardTunnel()
        secondary = tunnel.NewShadowsocksTunnel()
    }

    if e.caps.UseUTLS {
        obfs.EnableUTLS(e.caps.UTLSProfile)
    }
    if e.caps.FragSize > 0 {
        obfs.EnableFragmentation(e.caps.FragSize)
    }
    if e.caps.AIMorphEnabled {
        ai_morph.EnableTrafficMorphing()
    }

    // تحلیلگر تطبیقی DPI
    var dpiAnalyzer *dpi_analyzer.Analyzer
    if e.caps.DPIAdaptive {
        dpiAnalyzer = dpi_analyzer.NewAnalyzer()
        go func() {
            for {
                time.Sleep(5 * time.Second)
                bestProfile := rl_agent.GetRecommendedProfile()
                dpiAnalyzer.SetProfile(bestProfile)
                ai_morph.SetActiveProfile(bestProfile)
            }
        }()
    }

    active_shield.StartProbeShield()

    // نشست کوانتومی
    var pqcSession *pqc.QuantumSession
    if e.caps.PQCLevel > 0 {
        sess, err := pqc.NewQuantumSession(e.caps.PQCLevel)
        if err != nil {
            log.Printf("PQC init error: %v", err)
        } else {
            pqcSession = sess
            primary.SetQuantumSession(sess)
            secondary.SetQuantumSession(sess)
            sess.RegisterTunnel(primary)
            sess.StartKeyRollover()
        }
    }

    // تونل‌های مخفی اضطراری
    if e.caps.UseDNSTunnel {
        dnsTun := dns_tunnel.NewDNSTunnel("tunnel.example.com", pqcSession)
        dnsTun.Start()
    }
    if e.caps.UseICMPTunnel {
        icmpTun := icmp_tunnel.NewICMPTunnel()
        if icmpTun != nil {
            icmpTun.Start()
        }
    }

    // QRNG entropy injection
    if e.caps.UseQRNG {
        log.Println("QRNG entropy injection active")
    }

    // یادگیری تقویتی
    rl_agent.StartLearning(primary, secondary)

    // تولید ترافیک پوششی مبتنی بر GAN
    gan_generator.EnableGANMorphing(30 * time.Second)

    // لفافه‌ساز QUIC برای پنهان‌سازی پروتکل
    quic_masq.EnableQUICWrapper(primary)

    // ارکستراسیون خودکار پروتکل
    orch := dynamic_orchestra.NewOrchestrator(primary, secondary)
    orch.Start()

    return primary.Start()
}
