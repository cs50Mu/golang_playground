package main

import "fmt"

func main() {
	res, err := ipv4ToInt("192.0.2.235")
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %v\n", res)
}
