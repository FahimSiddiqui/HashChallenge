package external

type ICache interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
}
