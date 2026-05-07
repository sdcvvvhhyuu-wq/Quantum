package dns_tunnel

import (
    "log"
    "github.com/miekg/dns"
    "github.com/argotunnel/core/pqc"
)

type DNSTunnel struct {
    domain string
    qs     *pqc.QuantumSession
}

func NewDNSTunnel(domain string, qs *pqc.QuantumSession) *DNSTunnel {
    return &DNSTunnel{domain: domain, qs: qs}
}

func (t *DNSTunnel) Start() {
    dns.HandleFunc(".", t.handleDNS)
    go func() {
        server := &dns.Server{Addr: ":5353", Net: "udp"}
        log.Println("DNS tunnel listening on :5353")
        if err := server.ListenAndServe(); err != nil {
            log.Printf("DNS tunnel error: %v", err)
        }
    }()
}

func (t *DNSTunnel) handleDNS(w dns.ResponseWriter, r *dns.Msg) {
    if len(r.Question) == 0 { return }
    qname := r.Question[0].Name
    log.Printf("DNS tunnel query: %s", qname)
    msg := new(dns.Msg)
    msg.SetReply(r)
    rr, _ := dns.NewRR(". IN A 127.0.0.1")
    msg.Answer = append(msg.Answer, rr)
    w.WriteMsg(msg)
}
