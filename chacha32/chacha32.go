package chacha32

import (
	"encoding/binary"
)

func rotl32(v uint32, n uint) uint32 {
	return (v << n) | (v >> (32 - n))
}

func quarterRound(a, b, c, d *uint32) {
	*a += *b
	*d ^= *a
	*d = rotl32(*d, 16)

	*c += *d
	*b ^= *c
	*b = rotl32(*b, 12)

	*a += *b
	*d ^= *a
	*d = rotl32(*d, 8)

	*c += *d
	*b ^= *c
	*b = rotl32(*b, 7)
}

func chacha32Block(key [8]uint32, nonce [3]uint32, counter uint32) [16]uint32 {
	constants := [4]uint32{
		0x61707865, // "expa"
		0x3320646e, // "nd 3"
		0x79622d32, // "2-by"
		0x6b206574, // "te k"
	}

	var state [16]uint32
	state[0], state[1], state[2], state[3] = constants[0], constants[1], constants[2], constants[3]
	copy(state[4:12], key[:])
	state[12] = counter
	state[13], state[14], state[15] = nonce[0], nonce[1], nonce[2]

	working := state

	for i := 0; i < 32; i += 2 { // 16 pilnų raundų kaip C kode
		// Column rounds
		quarterRound(&working[0], &working[4], &working[8], &working[12])
		quarterRound(&working[1], &working[5], &working[9], &working[13])
		quarterRound(&working[2], &working[6], &working[10], &working[14])
		quarterRound(&working[3], &working[7], &working[11], &working[15])

		// Diagonal rounds
		quarterRound(&working[0], &working[5], &working[10], &working[15])
		quarterRound(&working[1], &working[6], &working[11], &working[12])
		quarterRound(&working[2], &working[7], &working[8], &working[13])
		quarterRound(&working[3], &working[4], &working[9], &working[14])
	}

	var output [16]uint32
	for i := 0; i < 16; i++ {
		output[i] = working[i] + state[i]
	}
	return output
}

func bytesToUint32sLE(b []byte, n int) [8]uint32 {
	var out [8]uint32
	for i := 0; i < n; i++ {
		out[i] = binary.LittleEndian.Uint32(b[i*4 : (i+1)*4])
	}
	return out
}

func bytesToUint32sLE3(b []byte) [3]uint32 {
	var out [3]uint32
	for i := 0; i < 3; i++ {
		out[i] = binary.LittleEndian.Uint32(b[i*4 : (i+1)*4])
	}
	return out
}

func Encrypt(key, nonce, input []byte) []byte {
	keyWords := bytesToUint32sLE(key, 8)
	nonceWords := bytesToUint32sLE3(nonce)
	counter := uint32(0)
	output := make([]byte, len(input))

	for i := 0; i < len(input); i += 64 {
		block := chacha32Block(keyWords, nonceWords, counter)
		counter++

		var keystream [64]byte
		for j := 0; j < 16; j++ {
			binary.LittleEndian.PutUint32(keystream[j*4:(j+1)*4], block[j])
		}

		chunk := 64
		if len(input)-i < 64 {
			chunk = len(input) - i
		}
		for j := 0; j < chunk; j++ {
			output[i+j] = input[i+j] ^ keystream[j]
		}
	}
	return output
}

func Decrypt(key, nonce, input []byte) []byte {
	// Identical to Encrypt
	return Encrypt(key, nonce, input)
}
