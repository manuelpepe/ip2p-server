package ipconv

import (
	"encoding/binary"
	"errors"
	"net"
)

func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func Ip2int(ip net.IP) (uint32, error) {
	ip4 := ip.To4()
	if ip == nil {
		return 0, WrongIPFormatError
	}
	return binary.BigEndian.Uint32(ip4), nil
}

var WrongIPFormatError = errors.New("Wrong IP format")
