package main

import "fmt"

func main() {

	fmt.Println("Hello world this is go lang program")

	//Array and slice usage
	var arraydeclearation = [8]int{1, 3, 5, 6, 7, 8, 9}

	fmt.Println("Data inside array :", arraydeclearation)

	slicedata := arraydeclearation[2:3]
	fmt.Println("Data inside slice: ", slicedata)

	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))
	slicedata = append(slicedata, 4)
	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))
	slicedata = append(slicedata, 4)
	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))

	slicedata = append(slicedata, 4)
	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))

	slicedata = append(slicedata, 4)
	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))
	slicedata = append(slicedata, 4)
	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))

	slicedata = append(slicedata, 4)
	fmt.Println(cap(slicedata))
	fmt.Println(len(slicedata))

}
