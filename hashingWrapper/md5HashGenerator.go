package hashingWrapper

import (
	"crypto/md5"
	"encoding/hex"
)

type md5HashGenerator struct {
}

func NewHashGenerator() *md5HashGenerator {
	return &md5HashGenerator{}
}

func (m md5HashGenerator) GenerateHash(bytes []byte) string {
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}
