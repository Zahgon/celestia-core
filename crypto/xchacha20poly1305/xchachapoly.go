// Package xchacha20poly1305 creates an AEAD using hchacha, chacha, and poly1305
// This allows for randomized nonces to be used in conjunction with chacha.
package xchacha20poly1305

import (
	"crypto/cipher"
)

// Implements crypto.AEAD
type xchacha20poly1305 struct {
	key [KeySize]byte
}

const (
	// KeySize is the size of the key used by this AEAD, in bytes.
	KeySize = 32
	// NonceSize is the size of the nonce used with this AEAD, in bytes.
	NonceSize = 24
	// TagSize is the size added from poly1305
	TagSize = 16
	// MaxPlaintextSize is the max size that can be passed into a single call of Seal
	MaxPlaintextSize = (1 << 38) - 64
	// MaxCiphertextSize is the max size that can be passed into a single call of Open,
	// this differs from plaintext size due to the tag
	MaxCiphertextSize = (1 << 38) - 48

	// sigma are constants used in xchacha.
	// Unrolled from a slice so that they can be inlined, as slices can't be constants.
	sigma0 = uint32(0x61707865)
	sigma1 = uint32(0x3320646e)
	sigma2 = uint32(0x79622d32)
	sigma3 = uint32(0x6b206574)
)

// New returns a new xchachapoly1305 AEAD
func New(key []byte) (cipher.AEAD, error) { _ = "STUB: not implemented"; return *new(cipher.AEAD), nil }

func (c *xchacha20poly1305) NonceSize() int { _ = "STUB: not implemented"; return 0 }

func (c *xchacha20poly1305) Overhead() int { _ = "STUB: not implemented"; return 0 }

func (c *xchacha20poly1305) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// This can't error because we always provide a correctly sized key

func (c *xchacha20poly1305) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This can't error because we always provide a correctly sized key

// HChaCha exported from
// https://github.com/aead/chacha20/blob/8b13a72661dae6e9e5dea04f344f0dc95ea29547/chacha/chacha_generic.go#L194
// TODO: Add support for the different assembly instructions used there.

// The MIT License (MIT)

// Copyright (c) 2016 Andreas Auernhammer

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// HChaCha20 generates 32 pseudo-random bytes from a 128 bit nonce and a 256 bit secret key.
// It can be used as a key-derivation-function (KDF).
func HChaCha20(out *[32]byte, nonce *[16]byte, key *[32]byte) { _ = "STUB: not implemented"; return }

func hChaCha20Generic(out *[32]byte, nonce *[16]byte, key *[32]byte) {
	_ = "STUB: not implemented"
	return
}
