package main

import (
	"hash"
	"hash/fnv"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	set BitSet
	n uint // no of elements in filter
	k uint // no of hash value
	hashFunctions[] hash.Hash64
}

type Filter interface {
	Add(key []byte)
	Check(key [] byte) bool
	hashValues(key []byte) []uint64
}

func (b *BloomFilter) Add(key []byte){
	hashes := b.hashValues(key)
	i := uint(0)
	for i < b.k {
		pos := uint(hashes[i])%b.n
		b.set.Set(pos)
		i += 1
	}
}

func (b *BloomFilter) Check(key [] byte) bool{
	testValues := b.hashValues(key)
	i := uint(0)
	for i < b.k {
		pos := uint(testValues[i]) % b.n
		if b.set.Check(pos) == false {
			return false
		}
		i += 1
	}
	return true
}

func (b *BloomFilter) hashValues(key []byte) []uint64 {
	var results []uint64
	for _, hashFunc := range b.hashFunctions{
		hashFunc.Write(key)
		results = append(results, hashFunc.Sum64()) //64 bit hashing
		hashFunc.Reset()
	}
	return results
}

func NewFilter(len uint) Filter{
	var arr = NewBitSet(len)
	return &BloomFilter{
		arr,
		len,
		3,
		[]hash.Hash64{fnv.New64(), fnv.New64a(),murmur3.New64()},
	}
}