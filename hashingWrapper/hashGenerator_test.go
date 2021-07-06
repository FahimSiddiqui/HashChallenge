package hashingWrapper

import (
	"crypto"
	"strings"
	"testing"
)

func TestHashGeneration(t *testing.T) {
	hashGenerator := GetHashGenerator(crypto.MD5)
	hashedOutput := hashGenerator.GenerateHash([]byte("GoodMorning"))
	if strings.Compare(hashedOutput, "deaead14fbeb3c7d4bf835032bb63543") != 0 {
		t.Error()
	}
	if strings.Compare(hashedOutput, "INVALID_HASH_OUTPUT") == 0 {
		t.Error()
	}
}

func TestHashFactoryObj(t *testing.T) {
	hashGenerator := GetHashGenerator(crypto.MD5)
	if hashGenerator == nil {
		t.Fail()
	}

	// If MD4 is passed in the input and its not implemented, it would simply return nil.
	hashGenerator = GetHashGenerator(crypto.MD4)
	if hashGenerator != nil {
		t.Fail()
	}
}
