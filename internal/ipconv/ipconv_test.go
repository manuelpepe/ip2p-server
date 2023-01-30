package ipconv

import (
    "testing"
	"net"
)

func TestInt2ip(t *testing.T) {
	got := Int2ip(400762451)
	want := net.ParseIP("23.227.38.83")
	if got.String() != want.String() {
		t.Errorf("Expected %s got %s", want, got)
	}
}

func TestIp2int(t *testing.T) {
	got := Ip2int(net.ParseIP("23.227.38.83"))
	want := 400762451
	if got != uint32(want) {
		t.Errorf("Expected %d got %d", want, got)
	}
}