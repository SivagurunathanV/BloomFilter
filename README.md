# BloomFilter ![Read the Docs (version)](https://img.shields.io/readthedocs/pip/stable.svg)

Implemented Bloom Filter using Go Language

![Bloom Filter](https://en.wikipedia.org/wiki/Bloom_filter) is a data structure which check whether an element exist in the set or not.

1. bitSet.go

    Implemented a bitSet data structure to store key 
 
2. BloomFilter.go

    Implemented two main functions Add(key) and Check(key)

    Add(key) -> add element to the bitSet using K hashing functions provided.

    Check(key) -> hash the key and check each of the position in bitset. 

    If every bit position is set, then key is already exists (False positive).

    If its not set in anyone of them, key does not exists (False negative).

    As the number of Key increases in the bitSet, False positive rate increases (as more no. of keys may get collide i.e hash function does not distribution uniformly). 
    If the bloom filter report False negative, then its valid.
  
3. BloomFilterTest.go

     SampleTesting of the BloomFilter by adding few elements and checking it.
   

Run the bloomFilter using the command below

```
  go run main/BloomFilterTest.go main/BloomFilter.go main/bitSet.go
```

# Future Work:

Implement the bitSet data-structure to become more thread safe for distributed system applications.
