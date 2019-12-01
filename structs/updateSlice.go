package main

import "fmt"

func main() {
	//send copy of slice
	mySlice := []string{"husnu", "software", "go", "lang"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "husnutapan"
}
