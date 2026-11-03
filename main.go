package main

import (
	"fmt"
	"sync"
)

const (
	_ = 1 << iota
	first
	second
	third
)

var wg = sync.WaitGroup{}

func main() {
	type human struct {
		age       int16
		firstName string
		lastName  string
	}
	masoud := human{age: 32, firstName: "Masoud", lastName: "Esmaeilian"}
	info := map[string]string{
		"age":       "32",
		"firstName": "Masoud",
		"lastName":  "Esmaeilian",
	}
	myArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	temp := 'A'
	mySlice := myArray[:]
	const myConst int = 42
	var message1 string
	message1 = "Hello1"
	var message2 = "Hello2"
	var number float32 = 34.8
	var message3 = "Hello3"
	message4 := "Hello4"
	fmt.Printf("%v, %T \n", message1, message1)
	fmt.Printf("%v, %T \n", temp, temp)
	fmt.Printf("%v, %T \n", message2, message2)
	fmt.Printf("%v, %T \n", message3, message3)
	fmt.Printf("%v, %T \n", number, number)
	fmt.Printf("%v, %T \n", message4, message4)
	fmt.Printf("%v, %T \n", myConst, myConst)
	fmt.Printf("%v, %T \n", first, first)
	fmt.Printf("%v, %T \n", second, second)
	fmt.Printf("%v, %T \n", third, third)
	fmt.Printf("value: %v, length: %v, capacity: %v, type: %T \n", myArray, len(myArray), cap(myArray), myArray)
	fmt.Printf("value: %v, length: %v, capacity: %v, type: %T \n", mySlice, len(mySlice), cap(mySlice), mySlice)
	mySlice1 := append(mySlice, []int{11, 12, 13, 14, 15}...)
	fmt.Printf("value: %v, length: %v, capacity: %v, type: %T \n", mySlice1, len(mySlice1), cap(mySlice1), mySlice1)
	fmt.Printf("%v, %T \n", info, info)
	fmt.Printf("%v, %T \n", masoud, masoud)
	wg.Add(1)
	go sayHi()
	wg.Wait()

	ch := make(chan [2]int, 50)

	wg.Add(2)

	go func(ch chan<- [2]int) {
		numbers := [1000000]int{}

		for i, number := range numbers {
			ch <- [2]int{i, number}
		}
		close(ch)
		wg.Done()
	}(ch)

	go func(ch <-chan [2]int) {
		for {
			number, ok := <-ch
			if !ok {
				break
			}

			fmt.Println(number)
		}
		wg.Done()
	}(ch)
	wg.Wait()
}

func sayHi() {
	fmt.Println("Hi")
	wg.Done()
}
