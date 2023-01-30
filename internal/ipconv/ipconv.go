package ipconv

import (
	"net"
	"encoding/binary"

	"math"
	"strconv"
)

func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func Ip2int(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip)
}


func DecToIP(dec_ip uint32) string {
	var out string
	for i := 3; i >= 0; i-- {
		bytes_part := dec_ip / uint32(math.Pow(256, float64(i))) % 256
		out += strconv.FormatUint(uint64(bytes_part), 10)
		if i != 0 {
			out += "."
		}
	}
	return out
}