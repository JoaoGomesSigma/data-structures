package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	bloomFilter := new(100)

	bloomFilter.add([]byte("hello"))
	bloomFilter.add([]byte("world"))
	bloomFilter.add([]byte("test"))
	bloomFilter.add([]byte("02389erbdsf"))

	assert.Equal(t, true, bloomFilter.check([]byte("hello")))
	assert.Equal(t, false, bloomFilter.check([]byte("hi")))

	// fmt.Println(bloomFilter.bitArray)

	// 0.1 % of false positives
	// 1 in 692
}
