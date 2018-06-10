package main

import "fmt"

func main()  {

	var bloomFilter = NewFilter(1024)
	test1 := [] byte("siva")
	test2 := [] byte("guru")
	test3 := [] byte("nathan")

	fmt.Println("Adding value to Filter ->",string(test1))
	bloomFilter.Add(test1)

	fmt.Println("Adding value to Filter ->",string(test2))
	bloomFilter.Add([] byte("guru"))

	fmt.Println("Adding value to Filter ->",string(test3))
	bloomFilter.Add([] byte("nathan"))

	pTest := [] byte("nathan")
	nTest := [] byte("sivagurunathan")
	nTest2 := [] byte("Welcome to this wonderful world") // checking big size byte

	fmt.Println("Checking value in the Filter ->",string(pTest))
	fmt.Println(bloomFilter.Check(pTest))

	fmt.Println("Checking value in the Filter ->",string(nTest))
	fmt.Println(bloomFilter.Check(nTest))

	fmt.Println("Checking value in the Filter ->",string(nTest2))
	fmt.Println(bloomFilter.Check(nTest2))


}
