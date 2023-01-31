package ipconv

import (
	"net"
	"testing"
)

func TestInt2ip(t *testing.T) {
	got := Int2ip(400762451)
	want := net.ParseIP("23.227.38.83")
	if got.String() != want.String() {
		t.Errorf("Expected %s got %s", want, got)
	}
}

func TestIp2int(t *testing.T) {
	got, err := Ip2int(net.ParseIP("23.227.38.83"))
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	want := 400762451
	if got != uint32(want) {
		t.Errorf("Expected %d got %d", want, got)
	}
}

func TestIp2intBadFormat(t *testing.T) {
	type TC struct {
		arg    string
		expRes uint32
		expErr error
	}

	var tests = []TC{
		{"23.227.38.83", 400762451, nil},
		{"23.1", 0, WrongIPFormatError},
		{"23.227.38.83.123", 0, WrongIPFormatError},
		{"0.0.0.0", 0, nil},
		{"0.0.0.1", 1, nil},
		{"0.0.1.0", 256, nil},
		{"0.1.0.0", 65536, nil},
		{"1.0.0.0", 16777216, nil},
		{"255.255.255.255", 4294967295, nil},
		{"255.255.255.256", 0, WrongIPFormatError},
	}
	for _, test := range tests {
		got, err := Ip2int(net.ParseIP(test.arg))
		if err != test.expErr {
			t.Errorf("%s: Expected error to be %v got %v", test.arg, test.expErr, err)
		}
		if got != test.expRes {
			t.Errorf("%s: Expected %d got %d", test.arg, test.expRes, got)
		}
	}
}
