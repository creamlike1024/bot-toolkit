package toolkit

import (
	"fmt"

	"github.com/go-ping/ping"
)

func TraceRoute(ip string) {
	//TODO
}

// 模拟 ping
func Ping(ip string) (string, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return "", err
	}
	pinger.Count = 3

	var result string
	pinger.OnRecv = func(pkt *ping.Packet) {
		result += fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		result += fmt.Sprintf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		result += fmt.Sprintf("\n--- %s ping statistics ---\n", stats.Addr)
		result += fmt.Sprintf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		result += fmt.Sprintf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	result += fmt.Sprintf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
	return result, nil
}

func Curl(url string) {
	//TODO
}
