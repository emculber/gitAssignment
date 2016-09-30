package main

import "fmt"

func main() {
	Ninety_Nine_Bottles_of_Beer()
}

func Ninety_Nine_Bottles_of_Beer() {
	for i := 99; i > 0; i-- {
		fmt.Println(i, "bottles of beer on the wall,", i, "bottles of beer")
		if (i - 1) != 0 {
			fmt.Println("take one down, pass it around,", i-1, "bottles of beer on the wall.")
		} else {
			fmt.Println("take one down, pass it around, no more bottles of beer on the wall.")
		}
	}
}
