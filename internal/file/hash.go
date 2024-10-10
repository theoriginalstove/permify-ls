package file

import (
	"crypto/sha256"
	"fmt"
)

type Hash [sha256.Size]byte

func HashOf(data []byte) Hash {
	return Hash(sha256.Sum256(data))
}

func (h Hash) String() string {
	return fmt.Sprintf("%64x", [sha256.Size]byte(h))
}

func (h *Hash) XORWith(h2 Hash) {
	for i := range h {
		h[i] ^= h2[i]
	}
}
