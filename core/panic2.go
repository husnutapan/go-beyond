package main

import "fmt"

func main() {
	launchIde()
	fmt.Println("Main func completed...")
}

func launchIde() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Getting an error when launching ide.")
		}
	}()

	fmt.Println("Ide Starting....")
	startIde()
	fmt.Println("Ide Started...")

}

func startIde() {
	panic("Getting error...")
}
