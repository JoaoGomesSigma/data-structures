package bloomfilter

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray []bool
}

func new(size int) BloomFilter {
	return BloomFilter{
		bitArray: make([]bool, size),
	}
}

func (s *BloomFilter) add(item []byte) {
	results := hashFunctions(item)

	for _, hash := range results {
		position := uint(hash) % uint(len(s.bitArray))
		s.bitArray[position] = true
	}

}

func (s *BloomFilter) check(item []byte) bool {
	results := hashFunctions(item)

	for _, hash := range results {
		position := uint(hash) % uint(len(s.bitArray))
		if !s.bitArray[position] {
			return false
		}
	}

	return true
}

func hashFunctions(item []byte) []uint64 {
	results := make([]uint64, 3)

	hash1 := sha256.New()
	hash1.Write(item)
	sha256Sum := hash1.Sum(nil)
	results[0] = binary.LittleEndian.Uint64(sha256Sum)
	hash1.Reset()

	hash2 := fnv.New64()
	hash2.Write(item)
	results[1] = hash2.Sum64()
	hash2.Reset()

	hash3 := md5.New()
	hash3.Write(item)
	md5Sum := hash3.Sum(nil)
	results[2] = binary.LittleEndian.Uint64(md5Sum)
	hash3.Reset()

	return results
}
