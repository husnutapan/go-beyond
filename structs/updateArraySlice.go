package main

import "fmt"

func main() {
	//send copy of array
	mySlice := [5]string{"husnu", "software", "go", "lang"}
	updateArraySlice(mySlice)
	fmt.Println(mySlice)
}

func updateArraySlice(s [5]string) {
	s[0] = "husnutapan"
}
