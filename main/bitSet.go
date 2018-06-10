package main

import "fmt"

type bitSet struct {
	arr []bool
	len uint
}

type BitSet interface {
	Set(pos uint) bool
	Check(pos uint) bool
	Clear(pos uint) bool
	Len() uint
	Print()
}

func (b *bitSet) Set(pos uint) bool {
	if b == nil || b.len < pos  || pos < 0{
		return false
	}
	b.arr[pos] = true
	return b.arr[pos] == true
}

func (b *bitSet) Check(pos uint) bool{
	if b == nil || b.len < pos || pos < 0{
		return false
	}
	return b.arr[pos]
}

func (b *bitSet) Clear(pos uint) bool{
	if b == nil || b.len < pos || pos < 0{
		return false
	}
	b.arr[pos] = false
	if b.arr[pos] == false {
		return true
	}
	return false
}

func (b bitSet) Print()  {
	for _, v := range b.arr{
		fmt.Print(v, "\t")
	}
}

func (b *bitSet) Len() uint {
	if b == nil {
		return 0
	}
	return b.len
}

func NewBitSet(len uint) BitSet {
	fmt.Println("setting bit len", len)
	return &bitSet{make([] bool, len), len}
}

//func main(){
//	var b = NewBitSet(10)
//	fmt.Println(b.Len())
//	b.Set(1)
//	b.Print()
//}
