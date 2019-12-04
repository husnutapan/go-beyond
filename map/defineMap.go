package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}
	colors2 := make(map[string]string)
	colors2["white"] = "#fffff"
	delete(colors, "red")
	fmt.Println(colors)
	fmt.Println(colors2)

	sendReference(colors)
	printAll(colors)
}

func printAll(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("key:%v,value:%v", color, hex)
	}
}

func sendReference(colors map[string]string) {
	colors["white"] = "#11111"
}
