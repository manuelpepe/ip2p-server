package ipconv

import (
	"encoding/binary"
	"net"
)

func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func Ip2int(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}
