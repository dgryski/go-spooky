package spooky

import (
	"testing"
)

type _Golden struct {
	in   string
	out1 uint64
	out2 uint64
}

// This table generated from the reference C++ code
var shortGolden = []_Golden{
	{"", 0x232706fc6bf50919, 0x8b72ee65b4e851c7},
	{"a", 0x1a108191a0bbc9bd, 0x754258f061412a92},
	{"ab", 0xf9dbb6ad202a090f, 0x9c7059b0dad5ae93},
	{"abc", 0x8aab15f77537c967, 0xc61367f8ca7811b0},
	{"abcd", 0x5c6db4e0725121b4, 0xed4d2a6bf05f6d02},
	{"abcde", 0x4d36bf2cf609ea58, 0x19b91b8a95e63aad},
	{"abcdef", 0xe58acc6ac1806d46, 0xe883303f848b6936},
	{"abcdefg", 0xe134fb62c64ba57e, 0x82370d1a277e05e1},
	{"abcdefgh", 0x101c8730a539eb6e, 0xa8f6b7fcf2cdf12e},
	{"abcdefghi", 0xc07b3e2b5b2b9088, 0x325a9a3d862222a3},
	{"abcdefghij", 0x386052dd535ad608, 0x70e7f39d49914037},
	{"Discard medicine more than two years old.", 0x9a2a8b03f065d989, 0x75d5e55d5d40fec5},
	{"He who has a shady past knows that nice guys finish last.", 0x3cfae0aeec123aec, 0xd7735cc02943acc7},
	{"I wouldn't marry him with a ten foot pole.", 0x5cbb087a111a27e3, 0xfea6b212c325888e},
	{"Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", 0xefe646e8b596c5a9, 0xdb63dd14e1c109b3},
	{"The days of the digital watch are numbered.  -Tom Stoppard", 0x9cb53539a011d4d0, 0xde43b8d544062cfe},
	{"Nepal premier won't resign.", 0x1c00acb8421bf55b, 0xc1153d81ce2f097f},
	{"For every action there is an equal and opposite government program.", 0xd2c7109575a957b5, 0x9e154405d06f0ba0},
	{"His money is twice tainted: 'taint yours and 'taint mine.", 0x51a1c6d0f021f09, 0xe3d3136e05eb3fbc},
	{"There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", 0x27d9177ee154640f, 0x8059a9f29745ad82},
	{"It's a tiny change to the code and not completely disgusting. - Bob Manchek", 0x9a6dee0bcd39c963, 0xdf667ddcb5800dea},
	{"size:  a.out:  bad magic", 0x8f7e2206aca8d3e0, 0xbd8bec5d6fc0c3a5},
	{"The major problem is with sendmail.  -Mark Horton", 0xfd9d590950522f31, 0x80406cf282d5b68a},
	{"Give me a rock, paper and scissors and I will move the world.  CCFestoon", 0xc3359ca69be20184, 0xff5c3bb7100ffd49},
	{"If the enemy is within range, then so are you.", 0xd3f4f61c514cf6ca, 0x5a6b22a5b786528},
	{"It's well we cannot hear the screams/That we create in others' dreams.", 0x13e371bb70e685b5, 0xb393f60c6d429307},
	{"You remind me of a TV show, but that's all right: I watch it anyway.", 0x713c48c45813353f, 0xe9f202f216136ed1},
	{"C is as portable as Stonehedge!!", 0x9c3d456bfc68f45d, 0xf928cfd4cd627bd1},
	{"Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", 0x57fde96671cf1fc1, 0x55dcd3066faebcb4},
	{"The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", 0x995fd7a42818a4d4, 0x781feff2e19c18ca},
	{"How can you write a big system without C++?  -Paul Glick", 0x1793448a145061e7, 0xb89f4daa42c9a030},
}

func TestShort(t *testing.T) {

	for i, g := range shortGolden {
		var h1, h2 uint64
		Short([]byte(g.in), &h1, &h2)

		if h1 != g.out1 || h2 != g.out2 {
			t.Errorf("Short %d failed = 0x%x 0x%x want 0x%x 0x%x", i, h1, h2, g.out1, g.out2)
		}
	}
}
