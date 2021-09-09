package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().UnixNano()
	fmt.Println("time now is", now)
}
