package hashingWrapper

type IHashGenerator interface {
	GenerateHash([]byte) string
}
