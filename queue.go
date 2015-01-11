package main

import (
	"fmt"
	"github.com/adeven/redismq"
)

func main() {
	testQueue := redismq.CreateQueue("localhost", "6379", "", 9, "goqueue")

	fmt.Println("Added payoad")
	testQueue.Put("testpayload")
	
}
