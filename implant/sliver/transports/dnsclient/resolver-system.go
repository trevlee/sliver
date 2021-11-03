package dnsclient

import (
	"net"
	"time"

	// {{if .Config.Debug}}
	"log"
	// {{end}}
)

// Abstraction on top of miekg/dns and net/dns
type DNSResolver interface {
	Address() string
	A(string) ([][]byte, time.Duration, error)
	// CName(string) (string, time.Duration, error)
}

func NewSystemResolver() DNSResolver {
	return &SystemResolver{}
}

type SystemResolver struct{}

func (r *SystemResolver) Address() string {
	return "system"
}

func (r *SystemResolver) A(domain string) ([][]byte, time.Duration, error) {
	// {{if .Config.Debug}}
	log.Printf("[dns] %s->A record of %s?", r.Address(), domain)
	// {{end}}
	started := time.Now()
	ips, err := net.LookupIP(domain)
	rtt := time.Since(started)
	if err != nil {
		return nil, rtt, err
	}
	var addrs [][]byte
	for _, ip := range ips {
		if ip.To4() != nil {
			addrs = append(addrs, ip.To4())
		}
		if ip.To16() != nil {
			addrs = append(addrs, ip.To16())
		}
	}
	return addrs, rtt, nil
}
