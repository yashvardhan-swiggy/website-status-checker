package main

import (
	"fmt"
	"websites-status-checker/router"
)

func main() {
	fmt.Println("Server Started...")
	router.HandleRequest()
}
