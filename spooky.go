// Package spooky implements Bob Jenkins' Spooky hash
// http://www.burtleburtle.net/bob/hash/spooky.html
// Public domain, like the original
package spooky

import (
	"encoding/binary"
	"hash"
)

// number of uint64's in internal state
const sc_numVars = 12

// size of the internal state
const sc_blockSize = sc_numVars * 8

// size of buffer of unhashed data, in bytes
const sc_bufSize = sc_blockSize

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

func Mix(data []uint64, s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11 *uint64) {
	*s0 += data[0]
	*s2 ^= *s10
	*s11 ^= *s0
	*s0 = rot64(*s0, 11)
	*s11 += *s1
	*s1 += data[1]
	*s3 ^= *s11
	*s0 ^= *s1
	*s1 = rot64(*s1, 32)
	*s0 += *s2
	*s2 += data[2]
	*s4 ^= *s0
	*s1 ^= *s2
	*s2 = rot64(*s2, 43)
	*s1 += *s3
	*s3 += data[3]
	*s5 ^= *s1
	*s2 ^= *s3
	*s3 = rot64(*s3, 31)
	*s2 += *s4
	*s4 += data[4]
	*s6 ^= *s2
	*s3 ^= *s4
	*s4 = rot64(*s4, 17)
	*s3 += *s5
	*s5 += data[5]
	*s7 ^= *s3
	*s4 ^= *s5
	*s5 = rot64(*s5, 28)
	*s4 += *s6
	*s6 += data[6]
	*s8 ^= *s4
	*s5 ^= *s6
	*s6 = rot64(*s6, 39)
	*s5 += *s7
	*s7 += data[7]
	*s9 ^= *s5
	*s6 ^= *s7
	*s7 = rot64(*s7, 57)
	*s6 += *s8
	*s8 += data[8]
	*s10 ^= *s6
	*s7 ^= *s8
	*s8 = rot64(*s8, 55)
	*s7 += *s9
	*s9 += data[9]
	*s11 ^= *s7
	*s8 ^= *s9
	*s9 = rot64(*s9, 54)
	*s8 += *s10
	*s10 += data[10]
	*s0 ^= *s8
	*s9 ^= *s10
	*s10 = rot64(*s10, 22)
	*s9 += *s11
	*s11 += data[11]
	*s1 ^= *s9
	*s10 ^= *s11
	*s11 = rot64(*s11, 46)
	*s10 += *s0
}

func endPartial(h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11 *uint64) {
	*h11 += *h1
	*h2 ^= *h11
	*h1 = rot64(*h1, 44)
	*h0 += *h2
	*h3 ^= *h0
	*h2 = rot64(*h2, 15)
	*h1 += *h3
	*h4 ^= *h1
	*h3 = rot64(*h3, 34)
	*h2 += *h4
	*h5 ^= *h2
	*h4 = rot64(*h4, 21)
	*h3 += *h5
	*h6 ^= *h3
	*h5 = rot64(*h5, 38)
	*h4 += *h6
	*h7 ^= *h4
	*h6 = rot64(*h6, 33)
	*h5 += *h7
	*h8 ^= *h5
	*h7 = rot64(*h7, 10)
	*h6 += *h8
	*h9 ^= *h6
	*h8 = rot64(*h8, 13)
	*h7 += *h9
	*h10 ^= *h7
	*h9 = rot64(*h9, 38)
	*h8 += *h10
	*h11 ^= *h8
	*h10 = rot64(*h10, 53)
	*h9 += *h11
	*h0 ^= *h9
	*h11 = rot64(*h11, 42)
	*h10 += *h0
	*h1 ^= *h10
	*h0 = rot64(*h0, 54)
}

func end(data []uint64, h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11 *uint64) {
	*h0 += data[0]
	*h1 += data[1]
	*h2 += data[2]
	*h3 += data[3]
	*h4 += data[4]
	*h5 += data[5]
	*h6 += data[6]
	*h7 += data[7]
	*h8 += data[8]
	*h9 += data[9]
	*h10 += data[10]
	*h11 += data[11]
	endPartial(h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11)
	endPartial(h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11)
	endPartial(h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11)
}

func Hash128(message []byte, hash1, hash2 *uint64) {

	length := len(message)

	if length < 2*sc_blockSize {
		Short(message, hash1, hash2)
		return
	}

	var h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11 uint64
	var buf [sc_numVars]uint64
	u := message

	h0 = *hash1
	h1 = *hash2
	h2 = sc_const
	h3 = *hash1
	h4 = *hash2
	h5 = sc_const
	h6 = *hash1
	h7 = *hash2
	h8 = sc_const
	h9 = *hash1
	h10 = *hash2
	h11 = sc_const

	// handle all whole sc_blockSize blocks of bytes
	for len(u) >= sc_blockSize {
		for i := 0; i < sc_numVars; i++ {
			buf[i] = binary.LittleEndian.Uint64(u)
			u = u[8:]
		}
		Mix(buf[:], &h0, &h1, &h2, &h3, &h4, &h5, &h6, &h7, &h8, &h9, &h10, &h11)
	}

	remainder := len(u)

	// reset everything in buf
	for i := 0; i < sc_numVars; i++ {
		buf[i] = 0
	}

	// put in the data we have left
	var bidx int
	for bidx = 0; len(u) >= 8; bidx++ {
		buf[bidx] = binary.LittleEndian.Uint64(u)
		u = u[8:]
	}

	// we now have <1 uint64 left
	var tmpbuf [8]uint8
	copy(tmpbuf[:], u)

	buf[bidx] = binary.LittleEndian.Uint64(tmpbuf[:])
	bidx++

	// the last byte
	buf[sc_numVars-1] += uint64(remainder) << 56

	// do some final mixing

	end(buf[:], &h0, &h1, &h2, &h3, &h4, &h5, &h6, &h7, &h8, &h9, &h10, &h11)
	*hash1 = h0
	*hash2 = h1
}

func Hash64(message []byte, hash1 uint64) uint64 {
	hash2 := uint64(0)
	Hash128(message, &hash1, &hash2)
	return hash1
}

func Hash32(message []byte, hash1 uint32) uint32 {
	h1 := uint64(hash1)
	h2 := uint64(0)
	Hash128(message, &h1, &h2)
	return uint32(h1)
}

type spooky struct {
	m_data       [sc_bufSize]byte   // unhashed data, for partial messages
	m_state      [sc_numVars]uint64 // internal state of the hash
	m_length     int
	m_remainder  uint8
	seed1, seed2 uint64
}

func New(seed1, seed2 uint64) hash.Hash {
	h := &spooky{}
	h.seed1 = seed1
	h.seed2 = seed2
	h.m_state[0] = seed1
	h.m_state[1] = seed2
	return h
}

func (s *spooky) Reset() {
	s.m_length = 0
	s.m_remainder = 0
	s.m_state[0] = s.seed1
	s.m_state[1] = s.seed2
}

func (s *spooky) BlockSize() int { return 96 }
func (s *spooky) Size() int      { return 16 }

func (s *spooky) Sum(b []byte) []byte {

	// init the variables
	if s.m_length < sc_bufSize {
		hash1 := s.m_state[0]
		hash2 := s.m_state[1]
		Short(s.m_data[:s.m_length], &hash1, &hash2)

		var b1 [16]byte
		binary.LittleEndian.PutUint64(b1[:], hash1)
		binary.LittleEndian.PutUint64(b1[8:], hash2)
		return append(b, b1[:]...)
	}

	data := s.m_data[:s.m_remainder]
	remainder := s.m_remainder

	h0 := s.m_state[0]
	h1 := s.m_state[1]
	h2 := s.m_state[2]
	h3 := s.m_state[3]
	h4 := s.m_state[4]
	h5 := s.m_state[5]
	h6 := s.m_state[6]
	h7 := s.m_state[7]
	h8 := s.m_state[8]
	h9 := s.m_state[9]
	h10 := s.m_state[10]
	h11 := s.m_state[11]

	// mix in the last partial block, and the length mod sc_blockSize

	var buf [sc_numVars]uint64

	// put in the data we have left
	var bidx int
	for bidx = 0; len(data) >= 8; bidx++ {
		buf[bidx] = binary.LittleEndian.Uint64(data)
		data = data[8:]
	}

	// we now have <1 uint64 left
	var tmpbuf [8]uint8
	copy(tmpbuf[:], data)

	buf[bidx] = binary.LittleEndian.Uint64(tmpbuf[:])
	bidx++

	// the last byte
	buf[sc_numVars-1] += uint64(remainder) << 56

	// do some final mixing

	end(buf[:], &h0, &h1, &h2, &h3, &h4, &h5, &h6, &h7, &h8, &h9, &h10, &h11)

	var b1 [16]byte
	binary.LittleEndian.PutUint64(b1[:], h0)
	binary.LittleEndian.PutUint64(b1[8:], h1)

	return append(b, b1[:]...)
}

func (s *spooky) Write(message []byte) (int, error) {
	length := len(message)

	var h0, h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11 uint64
	newLength := length + int(s.m_remainder)

	u := message

	// Is this message fragment too short?  If it is, stuff it away.
	if newLength < sc_bufSize {
		copy(s.m_data[s.m_remainder:], message)
		s.m_length += length
		s.m_remainder = uint8(newLength)
		return length, nil
	}

	// init the variables
	if s.m_length < sc_bufSize {
		h0 = s.m_state[0]
		h1 = s.m_state[1]
		h2 = sc_const
		h3 = s.m_state[0]
		h4 = s.m_state[1]
		h5 = sc_const
		h6 = s.m_state[0]
		h7 = s.m_state[1]
		h8 = sc_const
		h9 = s.m_state[0]
		h10 = s.m_state[1]
		h11 = sc_const
	} else {
		h0 = s.m_state[0]
		h1 = s.m_state[1]
		h2 = s.m_state[2]
		h3 = s.m_state[3]
		h4 = s.m_state[4]
		h5 = s.m_state[5]
		h6 = s.m_state[6]
		h7 = s.m_state[7]
		h8 = s.m_state[8]
		h9 = s.m_state[9]
		h10 = s.m_state[10]
		h11 = s.m_state[11]
	}
	s.m_length += length

	// if we've got anything stuffed away, use it now
	if s.m_remainder != 0 {
		prefix := sc_bufSize - s.m_remainder
		copy(s.m_data[s.m_remainder:], message)

		var buf [sc_numVars]uint64
		for i := 0; i < sc_numVars; i++ {
			buf[i] = binary.LittleEndian.Uint64(s.m_data[i*8:])
		}

		Mix(buf[:], &h0, &h1, &h2, &h3, &h4, &h5, &h6, &h7, &h8, &h9, &h10, &h11)

		u = message[prefix:]
		length -= int(prefix)
	} else {
		u = message
	}

	// handle all whole blocks of sc_blockSize bytes
	for len(u) >= sc_blockSize {

		var buf [sc_numVars]uint64
		for i := 0; i < sc_numVars; i++ {
			buf[i] = binary.LittleEndian.Uint64(u)
			u = u[8:]
		}
		Mix(buf[:], &h0, &h1, &h2, &h3, &h4, &h5, &h6, &h7, &h8, &h9, &h10, &h11)
	}

	// stuff away the last few bytes
	s.m_remainder = uint8(len(u))
	copy(s.m_data[:], u)

	// stuff away the variables
	s.m_state[0] = h0
	s.m_state[1] = h1
	s.m_state[2] = h2
	s.m_state[3] = h3
	s.m_state[4] = h4
	s.m_state[5] = h5
	s.m_state[6] = h6
	s.m_state[7] = h7
	s.m_state[8] = h8
	s.m_state[9] = h9
	s.m_state[10] = h10
	s.m_state[11] = h11

	return length, nil
}
