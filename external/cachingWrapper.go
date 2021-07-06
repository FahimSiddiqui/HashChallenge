package external

import (
	"HashChallenge/hashingWrapper"
	"HashChallenge/interfaces"
	"crypto"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetUrlResponse(url *url.URL) (interfaces.ResponseStruct, error) {
	// fmt.Println(time.Now())
	// time.Sleep(3 * time.Second)
	inputURL := url.String()
	if value, ok := GetCacheInstance().get(inputURL); ok {
		// Cache hit, return same URL from the map, else try to compute it
		return interfaces.ResponseStruct{
			InputUrl:     inputURL,
			HashedOutput: value.(string),
		}, nil
	}
	// in case its not available in cache.
	res, err := http.Get(inputURL)
	if err != nil {
		return interfaces.ResponseStruct{}, err
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return interfaces.ResponseStruct{}, readErr
	}
	// Compute md5 hash
	var hashGenerator hashingWrapper.IHashGenerator
	hashGenerator = hashingWrapper.GetHashGenerator(crypto.MD5)
	hashedOutput := hashGenerator.GenerateHash(body)

	// HashedOutput can be stored back in the in-memory cache.
	GetCacheInstance().put(inputURL, hashedOutput)
	return interfaces.ResponseStruct{
		InputUrl:     inputURL,
		HashedOutput: hashedOutput,
	}, nil
}
