package k

import (
	"hash"
)

type Hasher111 interface {
	hash.Hash

	Hash([]byte) []byte
}
