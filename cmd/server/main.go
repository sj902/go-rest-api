package main

import "fmt"

func Run() error {
	fmt.Println("starting the app")
	return nil
}

func main() {
	fmt.Println("go api")

	if err := Run(); err !=nil {
		fmt.Println(err)
	}
}
