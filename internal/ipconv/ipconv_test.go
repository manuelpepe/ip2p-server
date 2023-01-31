package ipconv

import (
	"net"
	"testing"
)

func TestInt2ip(t *testing.T) {
	type TC struct {
		arg    uint32
		expRes string
	}

	var tests = []TC{
		{400762451, "23.227.38.83"},
		{0, "0.0.0.0"},
		{1, "0.0.0.1"},
		{256, "0.0.1.0"},
		{65536, "0.1.0.0"},
		{16777216, "1.0.0.0"},
		{4294967295, "255.255.255.255"},
	}

	for _, test := range tests {
		got := Int2ip(test.arg)
		want := net.ParseIP(test.expRes)
		if got.String() != want.String() {
			t.Errorf("Expected %s got %s", want, got)
		}
	}
}

func TestIp2int(t *testing.T) {
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
