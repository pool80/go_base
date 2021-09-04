package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now()
	fmt.Println(timestamp.Format("02 Jan 06 15:04"))
}