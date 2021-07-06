package hashingWrapper

import (
	"crypto"
)

// Mimicing Factory pattern

// GetHashGenerator := Create hashGenerator based on hash type.
func GetHashGenerator(hashType crypto.Hash) IHashGenerator {
	var hashGenerator IHashGenerator
	switch hashType {
	case crypto.MD5:
		hashGenerator = NewHashGenerator()
		break
	default:
		hashGenerator = nil
	}

	return hashGenerator
}
