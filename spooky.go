// Package spooky implements Bob Jenkins' Spooky hash
// http://www.burtleburtle.net/bob/hash/spooky.html
// Public domain, like the original
package spooky

import (
	"encoding/binary"
)

const sc_const = uint64(0xdeadbeefdeadbeef)

func rot64(x uint64, k int) uint64 {
	return (x << uint(k)) | (x >> (64 - uint(k)))
}

func shortMix(h0, h1, h2, h3 *uint64) {
	*h2 = rot64(*h2, 50)
	*h2 += *h3
	*h0 ^= *h2
	*h3 = rot64(*h3, 52)
	*h3 += *h0
	*h1 ^= *h3
	*h0 = rot64(*h0, 30)
	*h0 += *h1
	*h2 ^= *h0
	*h1 = rot64(*h1, 41)
	*h1 += *h2
	*h3 ^= *h1
	*h2 = rot64(*h2, 54)
	*h2 += *h3
	*h0 ^= *h2
	*h3 = rot64(*h3, 48)
	*h3 += *h0
	*h1 ^= *h3
	*h0 = rot64(*h0, 38)
	*h0 += *h1
	*h2 ^= *h0
	*h1 = rot64(*h1, 37)
	*h1 += *h2
	*h3 ^= *h1
	*h2 = rot64(*h2, 62)
	*h2 += *h3
	*h0 ^= *h2
	*h3 = rot64(*h3, 34)
	*h3 += *h0
	*h1 ^= *h3
	*h0 = rot64(*h0, 5)
	*h0 += *h1
	*h2 ^= *h0
	*h1 = rot64(*h1, 36)
	*h1 += *h2
	*h3 ^= *h1
}

func shortEnd(h0, h1, h2, h3 *uint64) {
	*h3 ^= *h2
	*h2 = rot64(*h2, 15)
	*h3 += *h2
	*h0 ^= *h3
	*h3 = rot64(*h3, 52)
	*h0 += *h3
	*h1 ^= *h0
	*h0 = rot64(*h0, 26)
	*h1 += *h0
	*h2 ^= *h1
	*h1 = rot64(*h1, 51)
	*h2 += *h1
	*h3 ^= *h2
	*h2 = rot64(*h2, 28)
	*h3 += *h2
	*h0 ^= *h3
	*h3 = rot64(*h3, 9)
	*h0 += *h3
	*h1 ^= *h0
	*h0 = rot64(*h0, 47)
	*h1 += *h0
	*h2 ^= *h1
	*h1 = rot64(*h1, 54)
	*h2 += *h1
	*h3 ^= *h2
	*h2 = rot64(*h2, 32)
	*h3 += *h2
	*h0 ^= *h3
	*h3 = rot64(*h3, 25)
	*h0 += *h3
	*h1 ^= *h0
	*h0 = rot64(*h0, 63)
	*h1 += *h0
}

func Short(message []byte, hash1, hash2 *uint64) {

	u := message

	length := len(u)

	remainder := length % 32
	a := *hash1
	b := *hash2
	c := sc_const
	d := sc_const

	if length > 15 {

		// handle all complete sets of 32 bytes
		for len(u) >= 32 {
			c += binary.LittleEndian.Uint64(u)
			d += binary.LittleEndian.Uint64(u[8:])
			shortMix(&a, &b, &c, &d)
			a += binary.LittleEndian.Uint64(u[16:])
			b += binary.LittleEndian.Uint64(u[24:])
			u = u[32:]
		}

		//Handle the case of 16+ remaining bytes.
		if remainder >= 16 {
			c += binary.LittleEndian.Uint64(u)
			d += binary.LittleEndian.Uint64(u[8:])
			shortMix(&a, &b, &c, &d)
			u = u[16:]
			remainder -= 16
		}
	}

	// Handle the last 0..15 bytes, and its length
	d += uint64(length) << 56
	switch remainder {
	case 15:
		d += uint64(u[14]) << 48
		fallthrough
	case 14:
		d += uint64(u[13]) << 40
		fallthrough
	case 13:
		d += uint64(u[12]) << 32
		fallthrough
	case 12:
		d += uint64(binary.LittleEndian.Uint32(u[8:]))
		c += binary.LittleEndian.Uint64(u[0:])
		break
	case 11:
		d += uint64(u[10]) << 16
		fallthrough
	case 10:
		d += uint64(u[9]) << 8
		fallthrough
	case 9:
		d += uint64(u[8])
		fallthrough
	case 8:
		c += binary.LittleEndian.Uint64(u[0:])
		break
	case 7:
		c += uint64(u[6]) << 48
		fallthrough
	case 6:
		c += uint64(u[5]) << 40
		fallthrough
	case 5:
		c += uint64(u[4]) << 32
		fallthrough
	case 4:
		c += uint64(binary.LittleEndian.Uint32(u[0:]))
		break
	case 3:
		c += uint64(u[2]) << 16
		fallthrough
	case 2:
		c += uint64(u[1]) << 8
		fallthrough
	case 1:
		c += uint64(u[0])
		break
	case 0:
		c += sc_const
		d += sc_const
	}
	shortEnd(&a, &b, &c, &d)
	*hash1 = a
	*hash2 = b
}
