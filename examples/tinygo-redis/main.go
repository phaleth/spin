package main

import (
	"fmt"

	"github.com/fermyon/spin/sdk/go/redis"
)

func init() {
	// redis.Handle() must be called in the init() function.
	redis.Handle(func(payload []byte) error {
		fmt.Println("Paylod::::")
		fmt.Println(string(payload))
		return nil
	})
}

// main functiion must be included for the compiler but is not executed.
func main() {}
